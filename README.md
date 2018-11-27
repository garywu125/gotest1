
Source version Control workflow:

1.建立遠端與近端的Source Code respository via Git
 - create remote source respository:
    - login GitHub.com
    - create a new repositoy 

- copy create local repository  script from github.com

- create local source repository:
   - execute local repository creation script

2. Interative commit/push process to keep remote/local repository sync;

3. add Tag release for the final commit to mark a milestone and sync tag update the remote repository:
   
    - git tag -a <tag-name> -m "message" : add new tag
    - git push origin <tag-name>

    - git tag :show all tags
    - git describe --tag : show current tag        
    - git checkout <tag-name> : switch tag release

4. build the local tag release to local specific directory: 
    - tag checkout <tag-name>    
    - using vscode differnet shell script to build differnt local tag release:
     - powershell: 
        - batch : PowerShell.exe -Command  "./release.ps1"
        - step by step process :
         $Env:GOOS = "linux"; $Env:GOARCH = "amd64";go build -o ../release/test1/v0.1.1/test1_linux_amd64
         $Env:GOOS = "windows"; $Env:GOARCH = "amd64";go build -o ../release/test1/v0.1.1/test1_windows_amd64.exe

     - git shell: 
         - batch process : git sh release.sh
         - step by step process :
         GOOS=linux GOARCH=amd64 go build -o ../release/test1/v0.1.1/test1_linux_amd64
         GOOS=windows GOARCH=amd64 go build -o ../release/test1/v0.1.1/test1_windows_amd64.exe

 5. 手動 publish local tage release build to be a remote repository release in github.com
     - draft a tage release on remote repository
     - upload local tage release build files to draft tag release
     - release the draft tag release on remote repository     