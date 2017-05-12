
func startServer(address string, sv *supervisor.Supervisor) (*grpc.Server, error) {
  l, err := net.Listen(defaultListenType, address)
  if err != nil {
    return nil, err
  }
  
  s := grpc.NewServer()
  types.RegisterAPIServer(s, server.NewServer(sv)
  go func() {
    if err := s.Serve(l); err != nil {
      logrus.WithField("error", err).Fatal("containerd: serve grpc")
    }
  }()
  return s, nil
}