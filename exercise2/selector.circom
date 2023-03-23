pragma circom 2.1.4;

include "circomlib/poseidon.circom";
// include "https://github.com/0xPARC/circom-secp256k1/blob/master/circuits/bigint.circom";

template Selector(nChoices) {
    signal input in[nChoices];
    signal input index;
    signal output out;

    signal d;

    d <-- nChoices < (index+1) ? 0 : in[index];
    
    out <== d;
}

component main = Selector(5);

/* INPUT = {
    "in": ["1", "2", "3", "4", "5"],
    "index": "4"
} */