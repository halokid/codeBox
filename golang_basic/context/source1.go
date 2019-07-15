// propagateCancel arranges for child to be canceled when parent is.
func propagateCancel(parent Context, child canceler) {
  if parent.Done() == nil {
    return // parent is never canceled
  }
  if p, ok := parentCancelCtx(parent); ok {
    p.mu.Lock()
    if p.err != nil {
      // parent has already been canceled
      child.cancel(false, p.err)
    } else {
      if p.children == nil {
        p.children = make(map[canceler]struct{})
      }
      p.children[child] = struct{}{}
    }
    p.mu.Unlock()
  } else {
    go func() {
      select {
      case <-parent.Done():
        child.cancel(false, parent.Err())
      case <-child.Done():
      }
    }()
  }
}

