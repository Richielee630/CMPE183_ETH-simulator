pragma solidity ^0.4.23;


contract Store {
    
    struct Trx {
        string msg;
    }
 
    mapping(uint256 => Trx ) public nonceToTrxs ; 


    function setItem(uint256 _nonce,string _msg) public  {

        nonceToTrxs[_nonce]=Trx(_msg);
    }
    
}

