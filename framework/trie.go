package framework

import (
	"errors"
	"github.com/yujian0213/self-web/framework/gin"
	"strings"
)
//代表树结构
type Tree struct {
	root *node
}
//节点
type node struct {
	isLast bool
	segment string
	handlers []gin.HandlerFunc
	childs []*node
	parent *node
}
func newNode() *node {
	return &node{
		isLast:  false,
		segment: "",
		childs:  []*node{},
	}
}
func NewTree() *Tree  {
	return &Tree{root: newNode()}
}
//判断一个segment是否是通用segment，即以：开头
func isWildSegment(segment string) bool  {
	return strings.HasPrefix(segment,":")
}
//过滤下一层满足segment规则的子节点
func (n *node) filterChildNodes(segment string) []*node {
	if len(n.childs) == 0 {
		return nil
	}
	//如果segment是通配符，该segment所有子节点都满足需求
	if isWildSegment(segment) {
		return n.childs
	}
	nodes := make([]*node,0,len(n.childs))
	//过滤所有下一层子节点
	for _, cnode := range n.childs {
		if isWildSegment(cnode.segment) {
			nodes = append(nodes,cnode)
		}else if cnode.segment == segment {
			nodes = append(nodes,cnode)
		}
	}
	return nodes
}
//判断路由是否已经在节点的所有子节点树中存在了
func (n *node) matchNode(uri string) *node {
	segments := strings.SplitN(uri,"/",2)
	segment := segments[0]
	if !isWildSegment(segment) {
		segment = strings.ToUpper(segment)
	}
	cnodes := n.filterChildNodes(segment)
	// 如果当前子节点没有一个符合，那么说明这个uri一定是之前不存在, 直接返回nil
	if cnodes == nil || len(cnodes) == 0 {
		return nil
	}
	// 如果只有一个segment，则是最后一个标记
	if len(segments) == 1 {
		// 如果segment已经是最后一个节点，判断这些cnode是否有isLast标志
		for _, tn := range cnodes {
			if tn.isLast {
				return tn
			}
		}

		// 都不是最后一个节点
		return nil
	}

	// 如果有2个segment, 递归每个子节点继续进行查找
	for _, tn := range cnodes {
		tnMatch := tn.matchNode(segments[1])
		if tnMatch != nil {
			return tnMatch
		}
	}
	return nil
}
//增加路由节点
func (tree *Tree) AddRouter(uri string, handlers []gin.HandlerFunc) error {
	n := tree.root
	if n.matchNode(uri) != nil {
		return errors.New("route exist: " + uri)
	}

	segments := strings.Split(uri, "/")
	// 对每个segment
	for index, segment := range segments {

		// 最终进入node segment的字段
		if !isWildSegment(segment) {
			segment = strings.ToUpper(segment)
		}
		isLast := index == len(segments)-1

		var objNode *node // 标记是否有合适的子节点

		childNodes := n.filterChildNodes(segment)
		// 如果有匹配的子节点
		if len(childNodes) > 0 {
			// 如果有segment相同的子节点，则选择这个子节点
			for _, cnode := range childNodes {
				if cnode.segment == segment {
					objNode = cnode
					break
				}
			}
		}

		if objNode == nil {
			// 创建一个当前node的节点
			cnode := newNode()
			cnode.segment = segment
			if isLast {
				cnode.isLast = true
				cnode.handlers = handlers
			}
			// 父节点指针修改
			cnode.parent = n
			n.childs = append(n.childs, cnode)
			objNode = cnode
		}

		n = objNode
	}

	return nil
}

// 匹配uri获取handler
func (tree *Tree) FindHandler(uri string) []gin.HandlerFunc {
	matchNode := tree.root.matchNode(uri)
	if matchNode == nil {
		return nil
	}
	return matchNode.handlers
}
// 将uri解析为params
func (n *node) parseParamsFromEndNode(uri string) map[string]string {
	ret := map[string]string{}
	segments := strings.Split(uri, "/")
	cnt := len(segments)
	cur := n
	for i := cnt - 1; i >= 0; i-- {
		if cur.segment == "" {
			break
		}
		// 如果是通配符节点
		if isWildSegment(cur.segment) {
			// 设置params
			ret[cur.segment[1:]] = segments[i]
		}
		cur = cur.parent
	}
	return ret
}