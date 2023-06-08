import rpa as r

if __name__ =='__main__':
  # r.init()
  # r.url('https://duckduckgo.com')
  # r.type('//*[@name="q"]', 'decentralisation[enter]')
  # r.wait()  # ensure results are fully loaded
  # r.snap('page', 'results.png')
  # r.close()

  # mouse
  r.init(visual_automation=True)
  r.type(600, 300, 'neo kobe city')
  r.click(900, 300)
  r.snap('page.png', 'results.png')
  r.hover('button_to_drag.png')
  r.mouse('down')
  r.hover(r.mouse_x() + 300, r.mouse_y())
  r.mouse('up')
  r.close()
