// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * SSH proxy local CA.
 *
 * ## Import
 *
 * FirewallSsh LocalCa can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:firewallssh/localca:Localca labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:firewallssh/localca:Localca labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Localca extends pulumi.CustomResource {
    /**
     * Get an existing Localca resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LocalcaState, opts?: pulumi.CustomResourceOptions): Localca {
        return new Localca(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:firewallssh/localca:Localca';

    /**
     * Returns true if the given object is an instance of Localca.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Localca {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Localca.__pulumiType;
    }

    /**
     * SSH proxy local CA name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Password for SSH private key.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * SSH proxy private key, encrypted with a password.
     */
    public readonly privateKey!: pulumi.Output<string>;
    /**
     * SSH proxy public key.
     */
    public readonly publicKey!: pulumi.Output<string>;
    /**
     * SSH proxy local CA source type. Valid values: `built-in`, `user`.
     */
    public readonly source!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Localca resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LocalcaArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LocalcaArgs | LocalcaState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LocalcaState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["privateKey"] = state ? state.privateKey : undefined;
            resourceInputs["publicKey"] = state ? state.publicKey : undefined;
            resourceInputs["source"] = state ? state.source : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as LocalcaArgs | undefined;
            if ((!args || args.privateKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'privateKey'");
            }
            if ((!args || args.publicKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'publicKey'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["privateKey"] = args?.privateKey ? pulumi.secret(args.privateKey) : undefined;
            resourceInputs["publicKey"] = args?.publicKey ? pulumi.secret(args.publicKey) : undefined;
            resourceInputs["source"] = args ? args.source : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password", "privateKey", "publicKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Localca.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Localca resources.
 */
export interface LocalcaState {
    /**
     * SSH proxy local CA name.
     */
    name?: pulumi.Input<string>;
    /**
     * Password for SSH private key.
     */
    password?: pulumi.Input<string>;
    /**
     * SSH proxy private key, encrypted with a password.
     */
    privateKey?: pulumi.Input<string>;
    /**
     * SSH proxy public key.
     */
    publicKey?: pulumi.Input<string>;
    /**
     * SSH proxy local CA source type. Valid values: `built-in`, `user`.
     */
    source?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Localca resource.
 */
export interface LocalcaArgs {
    /**
     * SSH proxy local CA name.
     */
    name?: pulumi.Input<string>;
    /**
     * Password for SSH private key.
     */
    password?: pulumi.Input<string>;
    /**
     * SSH proxy private key, encrypted with a password.
     */
    privateKey: pulumi.Input<string>;
    /**
     * SSH proxy public key.
     */
    publicKey: pulumi.Input<string>;
    /**
     * SSH proxy local CA source type. Valid values: `built-in`, `user`.
     */
    source?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
