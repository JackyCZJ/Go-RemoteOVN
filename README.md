## GO-RemoteOVN

- Control ovn by restful api. 
- go to `handler/ovn/` see test file to know schema of data. 
- Edit `conf/config.yaml` for db info.
if not need to login for edit ovn , go to ``router/router.go`` to delete all ``middleware.AuthMiddleware()``.

----------------------
### usage

```bash
cd $GOPATH
git clone http://git.esix.com/jackyczj/go-restfulovn.git
make ca
go install
go build .
```
#### OVN Option
```bash
ovn-nbctl set-connection ptcp:PORT[:IP] //for remote connection
ovn-nbctl set-connection pssl:PORT[:IP] //for remote secure connection
```
 

#### Run Status
```bash
./admin.sh start
./admin.sh status
./admin.sh stop
./admin.sh restart
```
todo
====
logical switch
--
- [x] ls Add
- [x] ls del
- [x] ls get
- [x] ls list
- [x] lsp add
- [x] lsp del
- [x] lsp List
- [x] ext id add
- [x] ext id del
- [x] LSP set Address
- [x] LSP Set Port Security

ACL
--
- [x] ACL add
- [x] ACL del
- [x] ACL List

Address Set
--
- [x] AS get
- [x] AS update
- [x] AS list
- [x] AS add
- [x] AS del

Logical Router
--
- [x] lr add
- [x] lr get
- [x] lr del
- [x] lr list
- [x] lrp list 
- [x] lrp add
- [x] lrp del

Load Balancer
--
- [x] lb add
- [x] lb del
- [x] lb get
- [x] lb updates
- [x] ls Load balance add
- [x] ls lb del
- [x] ls lb list
- [x] lr lb add
- [x] lr lb del
- [x] lr lb list

QoS
--
- [x] QoS Add
- [x] QoS Del
- [x] Qos List

DHCP Options
--
- [x] DHCPOptions Add
- [x] DHCPOptions Del
- [ ] DHCPOptions Set  (write but not work)
- [x] DHCPOptions List

LSP DHCP V4 v6 Options
--
- [x] LSPv4Options Get
- [x] LSPv4Options Set
- [x] LSPv6Options Get (not work)
- [x] LSPv6Options Set (?)

NAT
--
- [x] LR NAT OPTION
