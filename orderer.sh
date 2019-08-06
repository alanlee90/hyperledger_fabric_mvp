ORDERER_GENERAL_LOGLEVEL=info \
ORDERER_GENERAL_LISTENADDRESS=0.0.0.0 \
ORDERER_GENERAL_GENESISMETHOD=file \
ORDERER_GENERAL_GENESISFILE=/root/mvp/crypto-config/ordererOrganizations/ordererorg/orderers/orderer2.ordererorg/genesis.block \
ORDERER_GENERAL_LOCALMSPID=OrdererOrgMSP \
ORDERER_GENERAL_LOCALMSPDIR=/root/mvp/crypto-config/ordererOrganizations/ordererorg/orderers/orderer2.ordererorg/msp \
ORDERER_GENERAL_TLS_ENABLED=false \
ORDERER_GENERAL_TLS_PRIVATEKEY=/root/mvp/crypto-config/ordererOrganizations/ordererorg/orderers/orderer2.ordererorg/tls/server.key \
ORDERER_GENERAL_TLS_CERTIFICATE=/root/mvp/crypto-config/ordererOrganizations/ordererorg/orderers/orderer2.ordererorg/tls/server.crt \
ORDERER_GENERAL_TLS_ROOTCAS=[/root/mvp/crypto-config/ordererOrganizations/ordererorg/orderers/orderer0.ordererorg/tls/ca.crt,/root/mvp/crypto-config/peerOrganizations/orderorg/peers/peer0.orderorg/tls/ca.crt,/root/mvp/crypto-config/peerOrganizations/contractororg/peers/peer2.contractororg/tls/ca.crt,/root/mvp/crypto-config/peerOrganizations/subcontractororg/peers/peer4.subcontractororg/tls/ca.crt,/root/mvp/crypto-config/peerOrganizations/dataproviderorg/peers/peer6.dataproviderorg/tls/ca.crt,/root/mvp/crypto-config/peerOrganizations/governanceorg/peers/peer7.governanceorg/tls/ca.crt] \
CONFIGTX_ORDERER_BATCHTIMEOUT=1s \
CONFIGTX_ORDERER_ORDERERTYPE=kafka \
CONFIGTX_ORDERER_KAFKA_BROKERS=[kafka-zookeeper:9092] \
orderer


# ORDERER_GENERAL_LISTENADDRESS=orderer0 \