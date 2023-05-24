date_time=$(date -d "2019-07-03 18:50:22" +"%A %b %e %T %Z %Y")
mkdir -p "tegar/$date_time"
echo "tegar/$date_time"

tegar/Wednesday Jul  3 18:50:22 WIB 2019
➜  ~ touch about_me.txt
➜  ~ rm abaout_me.txt
rm: cannot remove 'abaout_me.txt': No such file or directory
➜  ~ rm about_me.txt
➜  ~ sudo mkdir /about_me
➜  ~ cd ~
➜  ~ pwd
/root
➜  ~ cd ~
➜  ~ ls
' at Thu Apr 13 20:23:07 WIB 2023'   automate.sh.save   tegar
 automate.sh                         ohmyzsh
➜  ~ mkdir about_me
➜  ~ cd ~
➜  ~ ls
 about_me                            automate.sh        ohmyzsh
' at Thu Apr 13 20:23:07 WIB 2023'   automate.sh.save   tegar
➜  ~ cd path/ke/about_me
cd: no such file or directory: path/ke/about_me
➜  ~ cd path/about_me
cd: no such file or directory: path/about_me
➜  ~ cd ~/tegar/about_me
cd: no such file or directory: /root/tegar/about_me
➜  ~ cd ~
➜  ~ ls
about_me  automate.sh  automate.sh.save  ohmyzsh  tegar
➜  ~ mkdir personal
➜  ~ cd ~
➜  ~ ls
about_me  automate.sh  automate.sh.save  ohmyzsh  tegar
➜  ~ mkdir -p about_me/personal
➜  ~ cd
➜  ~ ls
about_me  automate.sh  automate.sh.save  ohmyzsh  tegar
➜  ~ cd personal
cd: no such file or directory: personal
➜  ~ cd ~/about_me
➜  about_me cd personal
➜  personal nano facebook.txt
➜  personal ls
facebook.txt
➜  personal cd~
zsh: command not found: cd~
➜  personal pwd
/root/about_me/personal
➜  personal cd ~/about_me
➜  about_me mkdir profesional
➜  about_me cd profesional
➜  profesional nano linkedin.txt
➜  profesional cd ~/about_me
➜  about_me mv about_me ~/tegar/
mv: cannot stat 'about_me': No such file or directory
➜  about_me mv about_me tegar/
mv: cannot stat 'about_me': No such file or directory
➜  about_me mv about_me tegar/
mv: cannot stat 'about_me': No such file or directory
➜  about_me ls
personal  profesional
➜  about_me cd ~
➜  ~ mv about_me tegar/
➜  ~ cd ~/tegar
➜  tegar ls
 about_me   date  'Wednesday Jul  3 18:50:22 WIB 2019'
➜  tegar mkdir my friends
➜  tegar cd ~/my friends
cd: string not in pwd: /root/my
➜  tegar mkdir my_friends
➜  tegar cd ~/my_friends
cd: no such file or directory: /root/my_friends
➜  tegar cd ~/my_friends
cd: no such file or directory: /root/my_friends
➜  tegar ls
 about_me   date   my_friends  'Wednesday Jul  3 18:50:22 WIB 2019'
➜  tegar cd my_friends
➜  my_friends nano list_of_my_friends.txt
➜  my_friends cd tegar
cd: no such file or directory: tegar
➜  my_friends cd ~/tegar
➜  tegar mkdir my_system_info
➜  tegar cd my_system_info
➜  my_system_info nano about_this_laptop.txt
➜  my_system_info nano internet_connection.txt
➜  my_system_info cd ~/tegar
➜  tegar ls
 about_me   my_friends      'Wednesday Jul  3 18:50:22 WIB 2019'
 date       my_system_info
➜  tegar tree "tegar at Wed Jul 3 18:50:22 WIB 2019"
zsh: command not found: tree
➜  tegar ls -r
'Wednesday Jul  3 18:50:22 WIB 2019'   my_friends   about_me
 my_system_info                        date
➜  tegar ls -R
.:
 about_me   my_friends      'Wednesday Jul  3 18:50:22 WIB 2019'
 date       my_system_info

./about_me:
personal  profesional

./about_me/personal:
facebook.txt

./about_me/profesional:
linkedin.txt

./date:

./my_friends:
list_of_my_friends.txt

./my_system_info:
about_this_laptop.txt  internet_connection.txt

'./Wednesday Jul  3 18:50:22 WIB 2019':
➜  tegar
