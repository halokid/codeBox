



process := &libcontainer.Process {
  Args:   []string{"/bin/bash"},
  Env:    []string{"PATH=/bin"},
  User:   "daemon",
  Stdin:  os.Stdin,
  Stdout: os.Stdout,
  Stderr: os.Stderr,
}


err := container.Start(process)

if err != nil {
  logruns.Fatal(err)
  container.Destory()
  return
}


//wait for the process to finish.
_, err := process.Wait()
if err != nil {
  logruns.Fatal(err)
}

//destory the container
container.Destory()

