<pre>
  _____                ._____.   .__                     _________      .__  _____  _____ 
  /  _  \   ____   _____|__\_ |__ |  |   ____            /   _____/_____ |__|/ ____\/ ____\
 /  /_\  \ /    \ /  ___/  || __ \|  | _/ __ \   ______  \_____  \\____ \|  \   __\\   __\ 
/    |    \   |  \\___ \|  || \_\ \  |_\  ___/  /_____/  /        \  |_> >  ||  |   |  |   
\____|__  /___|  /____  >__||___  /____/\___  >         /_______  /   __/|__||__|   |__|   
        \/     \/     \/        \/          \/                  \/|__|   
</pre>
# Ansible-Spiff

 Ansible-Spiff is an Ansible binary module that provides support for using [Spiff](https://github.com/cloudfoundry-incubator/spiff) deployment manifest builder from within Ansible.
 
 The module wraps the official [Spiff](https://github.com/cloudfoundry-incubator/spiff) CLI and makes the _*merge*_ and _*diff*_ operations usable from Ansible.
 
# How to Install
 
 For more information on binary modules and how to install please refer to [Ansible's official documentation](http://docs.ansible.com/ansible/).
 
# Usage
  
  Usage of Ansible-Spiff's _*merge*_ and _*diff*_ operations through Ansible, is demostatrated and tested using the [spiff-merge.yml](https://github.com/ipolyzos/ansible-spiff/tree/master/examples/spiff-merge.yml) and [spiff-diff.yml](https://github.com/ipolyzos/ansible-spiff/tree/master/examples/spiff-diff.yml) playbooks in the [examples](https://github.com/ipolyzos/ansible-spiff/tree/master/examples) folder.
 
# Developing Ansible-Spiff
 
 See the [contribution guidelines](https://github.com/ipolyzos/ansible-spiff/tree/master/CONTRIBUTING.md).
 