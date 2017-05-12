


// ----- API interface -------- 
type apiServer struct {
  sv *supervisor.Supervisor
}



// ------- factory method ------
func NewServer(sv *supervisor.Supervisor) types.APIServer {
  return &apiServer {
    sv: sv,
  }
}



// ----- services hanlders' implementation ----
func (s *apiServer) CreateContainer(ctx context.Context, c *types.CreateContainerRequest) (*types.CreateContainerResponse, error) {
    if c.BundlePath == "" {
        return nil, errors.New("empty bundle path")
    }
    e := &supervisor.StartTask{}
    e.ID = c.Id
...
}
func (s *apiServer) Signal(ctx context.Context, r *types.SignalRequest) (*types.SignalResponse, error) {
...
}

func (s *apiServer) AddProcess(ctx context.Context, r *types.AddProcessRequest) (*types.AddProcessResponse, error) {
...}
...








