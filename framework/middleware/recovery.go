package middleware

import "self-web/framework"

func Recovery() framework.ControllerHandler {
	return func(c *framework.Context) error {
		defer func() {
			if err := recover(); err != nil {
				_ = c.SetStatus(500).Json( err)
			}
		}()
		_ = c.Next()
		return nil
	}
}
