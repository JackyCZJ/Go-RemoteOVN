## GO-RemoteOVN

- Control ovn by restful api. 
- go to `handler/ovn/` see test file to know schema of data. 
- Edit `conf/config.yaml` for db info.
if not need to login for edit ovn , go to ``router/router.go`` to delete all ``middleware.AuthMiddleware()``.
- Base on [ebay/go-ovn](https://github.com/ebay/go-ovn) , use fork [JackyCZJ/go-ovn](https://www.github.com/jackyczj/go-ovn)

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
```bash

 lrp-set-enabled PORT STATE
                            set administrative state PORT
                            ('enabled' or 'disabled')
  lrp-get-enabled PORT      get administrative state PORT
                            ('enabled' or 'disabled')
```

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
- [x] DHCPOptions Set
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

Meter
--
```bash
  meter-add NAME ACTION RATE UNIT [BURST] #add a meter
  meter-del [NAME]          #remove meters
  meter-list                #print meters
```
- [x] meter add
- [x] meter del
- [x] meter list

LRP Gateway Chassis
--
```bash
  lrp-set-gateway-chassis PORT CHASSIS [PRIORITY]
                            set gateway chassis for port PORT
  lrp-del-gateway-chassis PORT CHASSIS
                            delete gateway chassis from port PORT
  lrp-get-gateway-chassis PORT
```
- [ ] LRP SET gateway chassis
- [ ] LRP del ⬆️
- [ ] LRP get ⬆️



Things coming 🤣
--
```bash
#Connection commands:
  get-connection             #print the connections
  del-connection             #delete the connections
  [--inactivity-probe=MSECS]
  set-connection TARGET...   #set the list of connections to TARGET...

#SSL commands:
  get-ssl                     #print the SSL configuration
  del-ssl                     #delete the SSL configuration
  set-ssl PRIV-KEY CERT CA-CERT [SSL-PROTOS [SSL-CIPHERS]] #set the SSL configuration
#Port group commands:
  pg-add PG [PORTS]           #Create port group PG with optional PORTS
  pg-set-ports PG PORTS       #Set PORTS on port group PG
  pg-del PG                   #Delete port group PG

#HA chassis group commands:
  ha-chassis-group-add GRP  #Create an HA chassis group GRP
  ha-chassis-group-del GRP  #Delete the HA chassis group GRP
  ha-chassis-group-list     #List the HA chassis groups
  ha-chassis-group-add-chassis #GRP CHASSIS [PRIORITY] Adds an HAchassis with optional PRIORITY to the HA chassis group GRP
  ha-chassis-group-del-chassis #GRP CHASSIS Deletes the HA chassisCHASSIS from the HA chassis group GRP
```