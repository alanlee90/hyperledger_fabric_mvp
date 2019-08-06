CORE_PEER_ENDORSER_ENABLED=true \
CORE_PEER_ADDRESS=msitpeer:7051 \
CORE_PEER_CHAINCODELISTENADDRESS=msitpeer:7052 \
CORE_PEER_ID=governanceorg-peer9 \
CORE_PEER_LOCALMSPID=DataProviderOrgMSP \
CORE_PEER_GOSSIP_EXTERNALENDPOINT=msitpeer:7051 \
CORE_PEER_GOSSIP_USELEADERELECTION=true \
CORE_PEER_GOSSIP_ORGLEADER=false \
CORE_PEER_TLS_ENABLED=false \
CORE_PEER_TLS_KEY_FILE=/root/mvp/crypto-config/peerOrganizations/orderorg/peers/peer9.governanceorg/tls/server.key \
CORE_PEER_TLS_CERT_FILE=/root/mvp/crypto-config/peerOrganizations/governanceorg/peers/peer9.governanceorg/tls/server.crt \
CORE_PEER_TLS_ROOTCERT_FILE=/root/mvp/crypto-config/peerOrganizations/governanceorg/peers/peer9.governanceorg/tls/ca.crt \
CORE_PEER_TLS_SERVERHOSTOVERRIDE=msitpeer \
CORE_PEER_MSPCONFIGPATH=/root/mvp/crypto-config/peerOrganizations/governanceorg/peers/peer9.governanceorg/msp \
peer node start