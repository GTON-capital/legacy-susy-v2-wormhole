"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    Object.defineProperty(o, k2, { enumerable: true, get: function() { return m[k]; } });
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (_) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
var yargs_1 = __importDefault(require("yargs"));
var hideBin = require('yargs/helpers').hideBin;
var bridge = __importStar(require("bridge"));
var elliptic = __importStar(require("elliptic"));
var ethers = __importStar(require("ethers"));
var token_bridge = __importStar(require("token-bridge"));
var web3s = __importStar(require("@solana/web3.js"));
var js_base64_1 = require("js-base64");
var ethers_contracts_1 = require("./src/ethers-contracts");
var terra_js_1 = require("@terra-money/terra.js");
var terra_js_2 = require("@terra-money/terra.js");
var web3_js_1 = require("@solana/web3.js");
var utils_1 = require("ethers/lib/utils");
var signAndEncodeVM = function (timestamp, nonce, emitterChainId, emitterAddress, sequence, data, signers, guardianSetIndex, consistencyLevel) {
    var body = [
        ethers.utils.defaultAbiCoder.encode(["uint32"], [timestamp]).substring(2 + (64 - 8)),
        ethers.utils.defaultAbiCoder.encode(["uint32"], [nonce]).substring(2 + (64 - 8)),
        ethers.utils.defaultAbiCoder.encode(["uint16"], [emitterChainId]).substring(2 + (64 - 4)),
        ethers.utils.defaultAbiCoder.encode(["bytes32"], [emitterAddress]).substring(2),
        ethers.utils.defaultAbiCoder.encode(["uint64"], [sequence]).substring(2 + (64 - 16)),
        ethers.utils.defaultAbiCoder.encode(["uint8"], [consistencyLevel]).substring(2 + (64 - 2)),
        data.substr(2)
    ];
    var hash = utils_1.solidityKeccak256(["bytes"], [utils_1.solidityKeccak256(["bytes"], ["0x" + body.join("")])]);
    var signatures = "";
    for (var i in signers) {
        var ec = new elliptic.ec("secp256k1");
        var key = ec.keyFromPrivate(signers[i]);
        var signature = key.sign(Buffer.from(hash.substr(2), "hex"), { canonical: true });
        var packSig = [
            ethers.utils.defaultAbiCoder.encode(["uint8"], [i]).substring(2 + (64 - 2)),
            zeroPadBytes(signature.r.toString(16), 32),
            zeroPadBytes(signature.s.toString(16), 32),
            ethers.utils.defaultAbiCoder.encode(["uint8"], [signature.recoveryParam]).substr(2 + (64 - 2)),
        ];
        signatures += packSig.join("");
    }
    var vm = [
        ethers.utils.defaultAbiCoder.encode(["uint8"], [1]).substring(2 + (64 - 2)),
        ethers.utils.defaultAbiCoder.encode(["uint32"], [guardianSetIndex]).substring(2 + (64 - 8)),
        ethers.utils.defaultAbiCoder.encode(["uint8"], [signers.length]).substring(2 + (64 - 2)),
        signatures,
        body.join("")
    ].join("");
    return vm;
};
function zeroPadBytes(value, length) {
    while (value.length < 2 * length) {
        value = "0" + value;
    }
    return value;
}
yargs_1.default(hideBin(process.argv))
    .command('generate_register_chain_vaa [chain_id] [contract_address]', 'create a VAA to register a chain (debug-only)', function (yargs) {
    return yargs
        .positional('chain_id', {
        describe: 'chain id to register',
        type: "number",
        required: true
    })
        .positional('contract_address', {
        describe: 'contract to register',
        type: "string",
        required: true
    });
}, function (argv) { return __awaiter(void 0, void 0, void 0, function () {
    var data, vm;
    return __generator(this, function (_a) {
        data = [
            "0x",
            "000000000000000000000000000000000000000000546f6b656e427269646765",
            "01",
            "0000",
            ethers.utils.defaultAbiCoder.encode(["uint16"], [argv.chain_id]).substring(2 + (64 - 4)),
            ethers.utils.defaultAbiCoder.encode(["bytes32"], [argv.contract_address]).substring(2),
        ].join('');
        vm = signAndEncodeVM(1, 1, 1, "0x0000000000000000000000000000000000000000000000000000000000000004", Math.floor(Math.random() * 100000000), data, [
            "cfb12303a19cde580bb4dd771639b0d26bc68353645571a8cff516ab2ee113a0"
        ], 0, 0);
        console.log(vm);
        return [2 /*return*/];
    });
}); })
    .command('terra execute_governance_vaa [vaa]', 'execute a governance VAA on Terra', function (yargs) {
    return yargs
        .positional('vaa', {
        describe: 'vaa to post',
        type: "string",
        required: true
    })
        .option('rpc', {
        alias: 'u',
        type: 'string',
        description: 'URL of the Terra RPC',
        default: "http://localhost:1317"
    })
        .option('token_bridge', {
        alias: 't',
        type: 'string',
        description: 'Token Bridge address',
        default: "terra10pyejy66429refv3g35g2t7am0was7ya7kz2a4"
    })
        .option('chain_id', {
        alias: 'c',
        type: 'string',
        description: 'Chain ID',
        // Should be localterra in theory, however Terra Station will
        // assume columbus-4 when localterra is set, while our current
        // dev environment is based on columbus-4. Should change when
        // change ID within terra/devnet/config/genesis.json is also
        // changed.
        default: 'columbus-4'
    })
        .option('mnemonic', {
        alias: 'm',
        type: 'string',
        description: 'Wallet Mnemonic',
        default: 'notice oak worry limit wrap speak medal online prefer cluster roof addict wrist behave treat actual wasp year salad speed social layer crew genius',
    });
}, function (argv) { return __awaiter(void 0, void 0, void 0, function () {
    var terra, wallet, vaa, transaction;
    return __generator(this, function (_a) {
        terra = new terra_js_1.LCDClient({
            URL: argv.rpc,
            chainID: argv.chain_id,
        });
        wallet = terra.wallet(new terra_js_1.MnemonicKey({
            mnemonic: argv.mnemonic
        }));
        vaa = Buffer.from(argv.vaa, "hex");
        transaction = new terra_js_2.MsgExecuteContract(wallet.key.accAddress, argv.token_bridge, {
            submit_vaa: {
                data: js_base64_1.fromUint8Array(vaa)
            },
        }, { uluna: 1000 });
        wallet
            .createAndSignTx({
            msgs: [transaction],
            memo: '',
        })
            .then(function (tx) { return terra.tx.broadcast(tx); })
            .then(function (result) {
            console.log(result);
            console.log("TX hash: " + result.txhash);
        });
        return [2 /*return*/];
    });
}); })
    .command('solana execute_governance_vaa [vaa]', 'execute a governance VAA on Solana', function (yargs) {
    return yargs
        .positional('vaa', {
        describe: 'vaa to post',
        type: "string",
        required: true
    })
        .option('rpc', {
        alias: 'u',
        type: 'string',
        description: 'URL of the Solana RPC',
        default: "http://localhost:8899"
    })
        .option('bridge', {
        alias: 'b',
        type: 'string',
        description: 'Bridge address',
        default: "Bridge1p5gheXUvJ6jGWGeCsgPKgnE3YgdGKRVCMY9o"
    })
        .option('token_bridge', {
        alias: 't',
        type: 'string',
        description: 'Token Bridge address',
        default: "B6RHG3mfcckmrYN1UhmJzyS1XX3fZKbkeUcpJe9Sy3FE"
    });
}, function (argv) { return __awaiter(void 0, void 0, void 0, function () {
    var connection, bridge_id, token_bridge_id, from, airdropSignature, vaa, parsed_vaa, ix, transaction, signature;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                connection = setupConnection(argv);
                bridge_id = new web3_js_1.PublicKey(argv.bridge);
                token_bridge_id = new web3_js_1.PublicKey(argv.token_bridge);
                from = web3s.Keypair.generate();
                return [4 /*yield*/, connection.requestAirdrop(from.publicKey, web3s.LAMPORTS_PER_SOL)];
            case 1:
                airdropSignature = _a.sent();
                return [4 /*yield*/, connection.confirmTransaction(airdropSignature)];
            case 2:
                _a.sent();
                vaa = Buffer.from(argv.vaa, "hex");
                return [4 /*yield*/, post_vaa(connection, bridge_id, from, vaa)];
            case 3:
                _a.sent();
                return [4 /*yield*/, bridge.parse_vaa(vaa)];
            case 4:
                parsed_vaa = _a.sent();
                switch (parsed_vaa.payload[32]) {
                    case 1:
                        console.log("Registering chain");
                        ix = token_bridge.register_chain_ix(token_bridge_id.toString(), bridge_id.toString(), from.publicKey.toString(), vaa);
                        break;
                    case 2:
                        console.log("Upgrading contract");
                        ix = token_bridge.upgrade_contract_ix(token_bridge_id.toString(), bridge_id.toString(), from.publicKey.toString(), from.publicKey.toString(), vaa);
                        break;
                    default:
                        throw new Error("unknown governance action");
                }
                transaction = new web3s.Transaction().add(ixFromRust(ix));
                return [4 /*yield*/, web3s.sendAndConfirmTransaction(connection, transaction, [from], {
                        skipPreflight: true
                    })];
            case 5:
                signature = _a.sent();
                console.log('SIGNATURE', signature);
                return [2 /*return*/];
        }
    });
}); })
    .command('eth execute_governance_vaa [vaa]', 'execute a governance VAA on Solana', function (yargs) {
    return yargs
        .positional('vaa', {
        describe: 'vaa to post',
        type: "string",
        required: true
    })
        .option('rpc', {
        alias: 'u',
        type: 'string',
        description: 'URL of the ETH RPC',
        default: "http://localhost:8545"
    })
        .option('token_bridge', {
        alias: 't',
        type: 'string',
        description: 'Token Bridge address',
        default: "0x0290FB167208Af455bB137780163b7B7a9a10C16"
    })
        .option('key', {
        alias: 'k',
        type: 'string',
        description: 'Private key of the wallet',
        default: "0x4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d"
    });
}, function (argv) { return __awaiter(void 0, void 0, void 0, function () {
    var provider, signer, t, tb, vaa, parsed_vaa, _a, _b, _c, _d, _e, _f, _g;
    return __generator(this, function (_h) {
        switch (_h.label) {
            case 0:
                provider = new ethers.providers.JsonRpcProvider(argv.rpc);
                signer = new ethers.Wallet(argv.key, provider);
                t = new ethers_contracts_1.BridgeImplementation__factory(signer);
                tb = t.attach(argv.token_bridge);
                vaa = Buffer.from(argv.vaa, "hex");
                return [4 /*yield*/, bridge.parse_vaa(vaa)];
            case 1:
                parsed_vaa = _h.sent();
                _a = parsed_vaa.payload[32];
                switch (_a) {
                    case 1: return [3 /*break*/, 2];
                    case 2: return [3 /*break*/, 4];
                }
                return [3 /*break*/, 6];
            case 2:
                console.log("Registering chain");
                _c = (_b = console).log;
                _d = "Hash: ";
                return [4 /*yield*/, tb.registerChain(vaa)];
            case 3:
                _c.apply(_b, [_d + (_h.sent()).hash]);
                return [3 /*break*/, 7];
            case 4:
                console.log("Upgrading contract");
                _f = (_e = console).log;
                _g = "Hash: ";
                return [4 /*yield*/, tb.upgrade(vaa)];
            case 5:
                _f.apply(_e, [_g + (_h.sent()).hash]);
                return [3 /*break*/, 7];
            case 6: throw new Error("unknown governance action");
            case 7: return [2 /*return*/];
        }
    });
}); })
    .argv;
function post_vaa(connection, bridge_id, payer, vaa) {
    return __awaiter(this, void 0, void 0, function () {
        var bridge_state, guardian_addr, acc, guardian_data, signature_set, txs, _i, txs_1, tx, ixs, transaction_1, ix, transaction, signature;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0: return [4 /*yield*/, get_bridge_state(connection, bridge_id)];
                case 1:
                    bridge_state = _a.sent();
                    guardian_addr = new web3_js_1.PublicKey(bridge.guardian_set_address(bridge_id.toString(), bridge_state.guardian_set_index));
                    return [4 /*yield*/, connection.getAccountInfo(guardian_addr)];
                case 2:
                    acc = _a.sent();
                    if ((acc === null || acc === void 0 ? void 0 : acc.data) === undefined) {
                        return [2 /*return*/];
                    }
                    guardian_data = bridge.parse_guardian_set(new Uint8Array(acc === null || acc === void 0 ? void 0 : acc.data));
                    signature_set = web3_js_1.Keypair.generate();
                    txs = bridge.verify_signatures_ix(bridge_id.toString(), payer.publicKey.toString(), bridge_state.guardian_set_index, guardian_data, signature_set.publicKey.toString(), vaa);
                    _i = 0, txs_1 = txs;
                    _a.label = 3;
                case 3:
                    if (!(_i < txs_1.length)) return [3 /*break*/, 6];
                    tx = txs_1[_i];
                    ixs = tx.map(function (v) {
                        return ixFromRust(v);
                    });
                    transaction_1 = new web3s.Transaction().add(ixs[0], ixs[1]);
                    // Sign transaction, broadcast, and confirm
                    return [4 /*yield*/, web3s.sendAndConfirmTransaction(connection, transaction_1, [payer, signature_set], {
                            skipPreflight: true
                        })];
                case 4:
                    // Sign transaction, broadcast, and confirm
                    _a.sent();
                    _a.label = 5;
                case 5:
                    _i++;
                    return [3 /*break*/, 3];
                case 6:
                    ix = ixFromRust(bridge.post_vaa_ix(bridge_id.toString(), payer.publicKey.toString(), signature_set.publicKey.toString(), vaa));
                    transaction = new web3s.Transaction().add(ix);
                    return [4 /*yield*/, web3s.sendAndConfirmTransaction(connection, transaction, [payer], {
                            skipPreflight: true
                        })];
                case 7:
                    signature = _a.sent();
                    console.log('SIGNATURE', signature);
                    return [2 /*return*/];
            }
        });
    });
}
function get_bridge_state(connection, bridge_id) {
    return __awaiter(this, void 0, void 0, function () {
        var bridge_state, acc;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    bridge_state = new web3_js_1.PublicKey(bridge.state_address(bridge_id.toString()));
                    return [4 /*yield*/, connection.getAccountInfo(bridge_state)];
                case 1:
                    acc = _a.sent();
                    if ((acc === null || acc === void 0 ? void 0 : acc.data) === undefined) {
                        throw new Error("bridge state not found");
                    }
                    return [2 /*return*/, bridge.parse_state(new Uint8Array(acc === null || acc === void 0 ? void 0 : acc.data))];
            }
        });
    });
}
function setupConnection(argv) {
    return new web3s.Connection(argv.rpc, 'confirmed');
}
function ixFromRust(data) {
    var keys = data.accounts.map(accountMetaFromRust);
    return new web3_js_1.TransactionInstruction({
        programId: new web3_js_1.PublicKey(data.program_id),
        data: Buffer.from(data.data),
        keys: keys,
    });
}
function accountMetaFromRust(meta) {
    return {
        pubkey: new web3_js_1.PublicKey(meta.pubkey),
        isSigner: meta.is_signer,
        isWritable: meta.is_writable,
    };
}
