# RSS Blog aggregator

The purpose of this project is to 
- Add RSS feeds from across the internet to be collected
- Store the collected posts in a PostgreSQL database
- Follow and unfollow RSS feeds that other users have added
- View summaries of the aggregated posts in the terminal, with a link to the full post


## Description

Ever want a place where you can keep up to date with you favorite blog, news sites, podcasts even? Well now you can!
This is a simple RSS aggregator where you cancontinuously fetch new posts, and store them in PostgreSQL. 
You can follow/unfollow feeds added by others and browse summarized posts in your terminal with links to the full articles.

## Getting Started

### Dependencies

* Go 1.25+
* PostgreSQL
macOS
```
brew install postgresql@15
```
linux/WSL (debian)
```
sudo apt update
sudo apt install postgresql postgresql-contrib
```

### Installing

* clone the repo
```
git clone https://github.com/AAlejandro8/RSS.git
```

### Executing program

* After cloning manually create a config file in your home directory ```~/.gatorconfig.json``` with the following content
```
{
"db_url": "postgres://username:password@host:port/database?sslmode=disable"
}
```
* Some commands
* login
```
go run . login <user>
```
* register (register before logging in)
```
go run . register <user>
```
* addfeed
```
go run . <feedName> <URL>
```
* follow/unfollow
```
go run . follow/unfollow <URL>
```
## Version History
* 0.1
    * Initial Release
