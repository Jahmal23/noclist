# gatedutil

Util library with various reuseable goodies

1. To run tests - 
    
    ./scripts/run_tests.sh
    
2. Release Process - 
    - A - Create a feature branch off of latest develop branch
    - B - Get that feature approved to merge to develop branch
    - C - Merge develop to master, with updated version tag 

3. Usage -
    - A - Provide Associated Project with Bitbucket SSH key to run on CircleCi
    - B - Update git config locally to use valid bitbucket credentials with access to this repo
        - Example: git config \
                     --global \
                     url."https://(username):CxAJMbgQHVVkQpmLbsMr@bitbucket.org".insteadOf \
                     "https://bitbucket.org"
                     
4. Gotchas - You may have to run this command before doing a "go get" to add a dependency - export GO111MODULE="on"
    https://stackoverflow.com/questions/56475313/how-to-fix-go-get-warning-modules-disabled-by-go111module-auto-in-gopath-src
        
        
