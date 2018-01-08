# phoenix_corelib

# glide https://github.com/Masterminds/glide
glide create                            # Start a new workspace
glide get github.com/Masterminds/cookoo # Get a package and add to glide.yaml
glide install                           # Install packages and dependencies
# work, work, work
go build                                # Go tools work normally
glide up                                # Update to newest versions of the package

#  下载单个包, 例如github.com/mattn/go-adodb
glide get --all-dependencies -s -v github.com/mattn/go-adodb