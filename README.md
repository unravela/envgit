# EnvGit
EnvGit is inspired by EnvDir and runs another command with env. variables loaded directly from 
file in git repository.

    envgit --url https://github.com/my/repo --file devel.env ./run_my_app.sh

## Motivation

I fall into situations, where I wanted to have centralized and versioned configuration but I 
wanted to avoid installation of systems like Consul or Etcd. I needed something very simple and lightweight. 

I wrote this tool as response for my need - having easy to setup, easy to manage centralized and 
versioned configuration.
