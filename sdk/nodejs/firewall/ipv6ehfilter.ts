// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure IPv6 extension header filter.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.firewall.Ipv6ehfilter("trname", {
 *     auth: "disable",
 *     destOpt: "disable",
 *     fragment: "disable",
 *     hopOpt: "disable",
 *     noNext: "disable",
 *     routing: "enable",
 * });
 * ```
 *
 * ## Import
 *
 * Firewall Ipv6EhFilter can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:firewall/ipv6ehfilter:Ipv6ehfilter labelname FirewallIpv6EhFilter
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:firewall/ipv6ehfilter:Ipv6ehfilter labelname FirewallIpv6EhFilter
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Ipv6ehfilter extends pulumi.CustomResource {
    /**
     * Get an existing Ipv6ehfilter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: Ipv6ehfilterState, opts?: pulumi.CustomResourceOptions): Ipv6ehfilter {
        return new Ipv6ehfilter(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:firewall/ipv6ehfilter:Ipv6ehfilter';

    /**
     * Returns true if the given object is an instance of Ipv6ehfilter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ipv6ehfilter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ipv6ehfilter.__pulumiType;
    }

    /**
     * Enable/disable blocking packets with the Authentication header (default = disable). Valid values: `enable`, `disable`.
     */
    public readonly auth!: pulumi.Output<string>;
    /**
     * Enable/disable blocking packets with Destination Options headers (default = disable). Valid values: `enable`, `disable`.
     */
    public readonly destOpt!: pulumi.Output<string>;
    /**
     * Enable/disable blocking packets with the Fragment header (default = disable). Valid values: `enable`, `disable`.
     */
    public readonly fragment!: pulumi.Output<string>;
    /**
     * Block specific Hop-by-Hop and/or Destination Option types (max. 7 types, each between 0 and 255, default = 0).
     */
    public readonly hdoptType!: pulumi.Output<number>;
    /**
     * Enable/disable blocking packets with the Hop-by-Hop Options header (default = disable). Valid values: `enable`, `disable`.
     */
    public readonly hopOpt!: pulumi.Output<string>;
    /**
     * Enable/disable blocking packets with the No Next header (default = disable) Valid values: `enable`, `disable`.
     */
    public readonly noNext!: pulumi.Output<string>;
    /**
     * Enable/disable blocking packets with Routing headers (default = enable). Valid values: `enable`, `disable`.
     */
    public readonly routing!: pulumi.Output<string>;
    /**
     * Block specific Routing header types (max. 7 types, each between 0 and 255, default =  0).
     */
    public readonly routingType!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Ipv6ehfilter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: Ipv6ehfilterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: Ipv6ehfilterArgs | Ipv6ehfilterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as Ipv6ehfilterState | undefined;
            resourceInputs["auth"] = state ? state.auth : undefined;
            resourceInputs["destOpt"] = state ? state.destOpt : undefined;
            resourceInputs["fragment"] = state ? state.fragment : undefined;
            resourceInputs["hdoptType"] = state ? state.hdoptType : undefined;
            resourceInputs["hopOpt"] = state ? state.hopOpt : undefined;
            resourceInputs["noNext"] = state ? state.noNext : undefined;
            resourceInputs["routing"] = state ? state.routing : undefined;
            resourceInputs["routingType"] = state ? state.routingType : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as Ipv6ehfilterArgs | undefined;
            resourceInputs["auth"] = args ? args.auth : undefined;
            resourceInputs["destOpt"] = args ? args.destOpt : undefined;
            resourceInputs["fragment"] = args ? args.fragment : undefined;
            resourceInputs["hdoptType"] = args ? args.hdoptType : undefined;
            resourceInputs["hopOpt"] = args ? args.hopOpt : undefined;
            resourceInputs["noNext"] = args ? args.noNext : undefined;
            resourceInputs["routing"] = args ? args.routing : undefined;
            resourceInputs["routingType"] = args ? args.routingType : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Ipv6ehfilter.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Ipv6ehfilter resources.
 */
export interface Ipv6ehfilterState {
    /**
     * Enable/disable blocking packets with the Authentication header (default = disable). Valid values: `enable`, `disable`.
     */
    auth?: pulumi.Input<string>;
    /**
     * Enable/disable blocking packets with Destination Options headers (default = disable). Valid values: `enable`, `disable`.
     */
    destOpt?: pulumi.Input<string>;
    /**
     * Enable/disable blocking packets with the Fragment header (default = disable). Valid values: `enable`, `disable`.
     */
    fragment?: pulumi.Input<string>;
    /**
     * Block specific Hop-by-Hop and/or Destination Option types (max. 7 types, each between 0 and 255, default = 0).
     */
    hdoptType?: pulumi.Input<number>;
    /**
     * Enable/disable blocking packets with the Hop-by-Hop Options header (default = disable). Valid values: `enable`, `disable`.
     */
    hopOpt?: pulumi.Input<string>;
    /**
     * Enable/disable blocking packets with the No Next header (default = disable) Valid values: `enable`, `disable`.
     */
    noNext?: pulumi.Input<string>;
    /**
     * Enable/disable blocking packets with Routing headers (default = enable). Valid values: `enable`, `disable`.
     */
    routing?: pulumi.Input<string>;
    /**
     * Block specific Routing header types (max. 7 types, each between 0 and 255, default =  0).
     */
    routingType?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Ipv6ehfilter resource.
 */
export interface Ipv6ehfilterArgs {
    /**
     * Enable/disable blocking packets with the Authentication header (default = disable). Valid values: `enable`, `disable`.
     */
    auth?: pulumi.Input<string>;
    /**
     * Enable/disable blocking packets with Destination Options headers (default = disable). Valid values: `enable`, `disable`.
     */
    destOpt?: pulumi.Input<string>;
    /**
     * Enable/disable blocking packets with the Fragment header (default = disable). Valid values: `enable`, `disable`.
     */
    fragment?: pulumi.Input<string>;
    /**
     * Block specific Hop-by-Hop and/or Destination Option types (max. 7 types, each between 0 and 255, default = 0).
     */
    hdoptType?: pulumi.Input<number>;
    /**
     * Enable/disable blocking packets with the Hop-by-Hop Options header (default = disable). Valid values: `enable`, `disable`.
     */
    hopOpt?: pulumi.Input<string>;
    /**
     * Enable/disable blocking packets with the No Next header (default = disable) Valid values: `enable`, `disable`.
     */
    noNext?: pulumi.Input<string>;
    /**
     * Enable/disable blocking packets with Routing headers (default = enable). Valid values: `enable`, `disable`.
     */
    routing?: pulumi.Input<string>;
    /**
     * Block specific Routing header types (max. 7 types, each between 0 and 255, default =  0).
     */
    routingType?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}