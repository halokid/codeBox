func checkErr(err error) {
  if err != nil {
    panic(err)
    // log.Fatal(err)
    // fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
    // os.Exit(1)
  }
}