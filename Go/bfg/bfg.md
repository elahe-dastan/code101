1. add code to git 
2. add passwords
3. add extra code again

And I suddenly understand that the passwords are on git and everybody can access them, so delete them fast 

4. delete the passwords and push 
5. check github and see they are deleted pooof

But wait, WTF, all the passwords are still there, you just have to go to the history part and all the passwords are just 
one commit behind. That's a shame, I should do something as soon as possible.

6. I use `bfg` to clean the history, let's see what it does

6.1 It's obvious that `bfg` needs the `Git database` of my repository to clean the passwords, so it needs the bare repo 
```shell script
git clone --mirror git@github.com:name/repo.git
```
the above command outputs `[repo-name].git`

6.2 download bfg jar 

6.3 run `bfg`
```shell script
java -jar bfg.jar --delete-files password.go repo.git
```
now the history should have been updated