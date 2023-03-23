pragma circom 2.1.4;

include "circomlib/poseidon.circom";
// include "https://github.com/0xPARC/circom-secp256k1/blob/master/circuits/bigint.circom";

template IsEqual() {
    signal input in[2];
    signal output out;
    signal diff;

    diff <-- (in[0] - in[1]) != 0 ? 1 : 0;
    out <== -diff + 1;
}

component main = IsEqual();

/* INPUT = {
    "in": ["1", "-1"]
} */
