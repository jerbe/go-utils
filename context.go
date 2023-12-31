package goutils

import "context"

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/5 19:07
  @describe :
*/

// ContextIsDone 判断上下文(Context)是否已经结束,完成
func ContextIsDone(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}
