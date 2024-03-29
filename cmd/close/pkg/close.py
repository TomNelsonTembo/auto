import psutil

#app_list = ['chrome', 'slack', 'nautilus' ,'code', 'gnome-terminal-server','gnome-system-monitor', 'firefox', 'deja-dup', 'gnome-software', 'blender', 'whatsapp-for-linux', 'gnome-text-editor']
app_list = ['gnome-terminal-server']
def close_app(app_name):
    running_apps=psutil.process_iter(['pid','name']) #returns names of running processes
    found=False
    for app in running_apps:
        sys_app=app.info.get('name').split('.')[0].lower()

        if sys_app in app_name.split() or app_name in sys_app:
            pid=app.info.get('pid') #returns PID of the given app if found running
            
            
            try: #deleting the app if asked app is running.(It raises error for some windows apps)
                app_pid = psutil.Process(pid)
                app_pid.terminate()
                found=True
            except: pass
            
        else: pass
    if not found:
        print(app_name+" not found running")
    else:
        print(app_name+'('+sys_app+')'+' closed')

for apps in app_list:
    close_app(apps)