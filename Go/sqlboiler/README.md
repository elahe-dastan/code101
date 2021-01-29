## What is $PATH
I ran `echo $PATH` and here's the result `/home/raha/anaconda3/bin:/home/raha/anaconda3/condabin:/home/linuxbrew/.linuxbrew/bin:/home/linuxbrew/.linuxbrew/sbin:/home/raha/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin`<br/>
If I put an executable in either one of these directories, I do not need to set the path to the executable / script, but 
I can run it by its name as a command. I can (permanently) add a directory to $PATH by adding `export PATH=$PATH:/path/to/dir`
to my `~/.profile` file.
## GOPATH GOROOT .bashrc