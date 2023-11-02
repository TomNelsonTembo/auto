import psutil

def display_apps():
    running_apps=psutil.process_iter(['pid','name'])
    found = False


    for app in running_apps:
        sys_app = app.info.get('name').split('.')[0].lower()
        pid =app.info.get('pid')

        

        try:
            print(pid, ' : ', sys_app )

        except: pass

display_apps()