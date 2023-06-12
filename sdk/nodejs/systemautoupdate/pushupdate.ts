// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure push updates. Applies to FortiOS Version `<= 7.0.0`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.systemautoupdate.Pushupdate("trname", {
 *     address: "0.0.0.0",
 *     override: "disable",
 *     port: 9443,
 *     status: "disable",
 * });
 * ```
 *
 * ## Import
 *
 * SystemAutoupdate PushUpdate can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:systemautoupdate/pushupdate:Pushupdate labelname SystemAutoupdatePushUpdate
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:systemautoupdate/pushupdate:Pushupdate labelname SystemAutoupdatePushUpdate
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Pushupdate extends pulumi.CustomResource {
    /**
     * Get an existing Pushupdate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PushupdateState, opts?: pulumi.CustomResourceOptions): Pushupdate {
        return new Pushupdate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:systemautoupdate/pushupdate:Pushupdate';

    /**
     * Returns true if the given object is an instance of Pushupdate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Pushupdate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Pushupdate.__pulumiType;
    }

    /**
     * Push update override server.
     */
    public readonly address!: pulumi.Output<string>;
    /**
     * Enable/disable push update override server. Valid values: `enable`, `disable`.
     */
    public readonly override!: pulumi.Output<string>;
    /**
     * Push update override port. (Do not overlap with other service ports)
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * Enable/disable push updates. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Pushupdate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PushupdateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PushupdateArgs | PushupdateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PushupdateState | undefined;
            resourceInputs["address"] = state ? state.address : undefined;
            resourceInputs["override"] = state ? state.override : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as PushupdateArgs | undefined;
            if ((!args || args.address === undefined) && !opts.urn) {
                throw new Error("Missing required property 'address'");
            }
            if ((!args || args.override === undefined) && !opts.urn) {
                throw new Error("Missing required property 'override'");
            }
            if ((!args || args.port === undefined) && !opts.urn) {
                throw new Error("Missing required property 'port'");
            }
            if ((!args || args.status === undefined) && !opts.urn) {
                throw new Error("Missing required property 'status'");
            }
            resourceInputs["address"] = args ? args.address : undefined;
            resourceInputs["override"] = args ? args.override : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Pushupdate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Pushupdate resources.
 */
export interface PushupdateState {
    /**
     * Push update override server.
     */
    address?: pulumi.Input<string>;
    /**
     * Enable/disable push update override server. Valid values: `enable`, `disable`.
     */
    override?: pulumi.Input<string>;
    /**
     * Push update override port. (Do not overlap with other service ports)
     */
    port?: pulumi.Input<number>;
    /**
     * Enable/disable push updates. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Pushupdate resource.
 */
export interface PushupdateArgs {
    /**
     * Push update override server.
     */
    address: pulumi.Input<string>;
    /**
     * Enable/disable push update override server. Valid values: `enable`, `disable`.
     */
    override: pulumi.Input<string>;
    /**
     * Push update override port. (Do not overlap with other service ports)
     */
    port: pulumi.Input<number>;
    /**
     * Enable/disable push updates. Valid values: `enable`, `disable`.
     */
    status: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
