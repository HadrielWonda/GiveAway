 cryptogen generate --config=./crypto-config.yaml


 mkdir channel-artifacts


 configtxgen -profile Genesis -outputBlock ./channel-artifacts/genesis.block -channelID genesischannel 
 configtxgen -profile giveawayChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID giveawaychannel
 configtxgen -profile giveawayChannel --outputAnchorPeersUpdate ./channel-artifacts/NGOOrgAnchorUpdate.tx -asOrg NGO -channelID giveawaychannel
 configtxgen -profile giveawayChannel --outputAnchorPeersUpdate ./channel-artifacts/VolOrgAnchorUpdate.tx -asOrg Volunteers -channelID giveawaychannel


 
 change file name under ca folder : ca-certp=.pem to cert.pem and private_sk to PRIVATE_KEY in bot ngo.com and vol.com


cli commands
for devpeer
export CORE_PEER_ADDRESS=devpeer:7051
export CORE_PEER_LOCALMSPID=NGOMSP
export CORE_PEER_MSPCONFIGPATH=/crypto-config/peerOrganizations/ngo.com/users/Admin@ngo.com/msp

for firstngo
export CORE_PEER_ADDRESS=firstngo:7051
export CORE_PEER_LOCALMSPID=NGOMSP
export CORE_PEER_MSPCONFIGPATH=/crypto-config/peerOrganizations/ngo.com/users/Admin@ngo.com/msp

for devvol
export CORE_PEER_ADDRESS=devvol:7051
export CORE_PEER_LOCALMSPID=VolunteersMSP
export CORE_PEER_MSPCONFIGPATH=/crypto-config/peerOrganizations/vol.com/users/Admin@vol.com/msp/

peer channel create -f channel.tx -o orderer:7050 -c giveawaychannel

installing chaincode 

peer chaincode install -n giveaway -v 0 -p giveawayChanincode
peer chaincode instantiate -n giveaway -v 0 -C giveawaychannel -c '{"args":[]}'



