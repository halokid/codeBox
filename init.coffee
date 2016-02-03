# Your init script
#
# Atom will evaluate this file each time a new window is opened. It is run
# after packages are loaded/activated and after the previous editor state
# has been restored.
#
# An example hack to log to the console when each text editor is saved.
#
# atom.workspace.observeTextEditors (editor) ->
#   editor.onDidSave ->
#     console.log "Saved! #{editor.getPath()}"

# place this snippet into init.coffee in ~/.atom directory

#这个只能是在切换窗口焦点的时候才更换title的文件名，有个屌用啊
#但是这个还是有点用的，这个可以在tab里面显示文件的全路径，暂时就这样先解决问题吧
#毕竟世事无完美
#start
atom.workspace.observeTextEditors (editor) ->
    if editor.getTitle() isnt "untitled"
        sp = editor.getPath().split('/')
        title = sp.slice(sp.length-2).join('/') # gives name of containing folder along with filename
        editor.getTitle = -> title
        editor.getLongTitle = -> title

for item in atom.workspace.getPaneItems()
    if item.emitter?
        item.emitter.emit "did-change-title", item.getTitle()
#end
