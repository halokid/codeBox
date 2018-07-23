package hello

func hello() string {
  words := []string {"hello", "func", "in", "package", "hello"}
  wl := len(words)

  sentence := ""
  for key, word := range words {
    sentence += word
    if key < wl - 1 {
      sentence += " "
    } else {
      sentence += "."
    }
  }
  return sentence
}


