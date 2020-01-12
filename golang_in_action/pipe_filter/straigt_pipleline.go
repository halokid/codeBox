package pipe_filter

type StraightPipeline struct {
  Name      string
  Filters    *[]Filter
}

func NewStraightPipeline(name string, filters ...Filter) *StraightPipeline {
  return &StraightPipeline{
    Name:       name,
    Filters:     &filters,
  }
}

func (f *StraightPipeline) Process(data Request) (Response, error) {
  var ret interface{}
  var err error
  for _, filter := range *f.Filters {
    ret, err = filter.Process(data)
    if err != nil {
      return nil, err
    }
    data = ret
  }
  return ret, nil
}