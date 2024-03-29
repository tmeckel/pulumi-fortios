// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure IP address management services. Applies to FortiOS Version `>= 7.0.2`.
 *
 * ## Import
 *
 * System Ipam can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:sys/ipam:Ipam labelname SystemIpam
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:sys/ipam:Ipam labelname SystemIpam
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Ipam extends pulumi.CustomResource {
    /**
     * Get an existing Ipam resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IpamState, opts?: pulumi.CustomResourceOptions): Ipam {
        return new Ipam(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:sys/ipam:Ipam';

    /**
     * Returns true if the given object is an instance of Ipam.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ipam {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ipam.__pulumiType;
    }

    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Configure IPAM pool subnet, Class A - Class B subnet.
     */
    public readonly poolSubnet!: pulumi.Output<string>;
    /**
     * Configure IPAM pools. The structure of `pools` block is documented below.
     */
    public readonly pools!: pulumi.Output<outputs.sys.IpamPool[] | undefined>;
    /**
     * Configure IPAM allocation rules. The structure of `rules` block is documented below.
     */
    public readonly rules!: pulumi.Output<outputs.sys.IpamRule[] | undefined>;
    /**
     * Configure the type of IPAM server to use. Valid values: `cloud`, `fabric-root`.
     */
    public readonly serverType!: pulumi.Output<string>;
    /**
     * Enable/disable IP address management services. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Ipam resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: IpamArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IpamArgs | IpamState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IpamState | undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["poolSubnet"] = state ? state.poolSubnet : undefined;
            resourceInputs["pools"] = state ? state.pools : undefined;
            resourceInputs["rules"] = state ? state.rules : undefined;
            resourceInputs["serverType"] = state ? state.serverType : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as IpamArgs | undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["poolSubnet"] = args ? args.poolSubnet : undefined;
            resourceInputs["pools"] = args ? args.pools : undefined;
            resourceInputs["rules"] = args ? args.rules : undefined;
            resourceInputs["serverType"] = args ? args.serverType : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Ipam.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Ipam resources.
 */
export interface IpamState {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Configure IPAM pool subnet, Class A - Class B subnet.
     */
    poolSubnet?: pulumi.Input<string>;
    /**
     * Configure IPAM pools. The structure of `pools` block is documented below.
     */
    pools?: pulumi.Input<pulumi.Input<inputs.sys.IpamPool>[]>;
    /**
     * Configure IPAM allocation rules. The structure of `rules` block is documented below.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.sys.IpamRule>[]>;
    /**
     * Configure the type of IPAM server to use. Valid values: `cloud`, `fabric-root`.
     */
    serverType?: pulumi.Input<string>;
    /**
     * Enable/disable IP address management services. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Ipam resource.
 */
export interface IpamArgs {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Configure IPAM pool subnet, Class A - Class B subnet.
     */
    poolSubnet?: pulumi.Input<string>;
    /**
     * Configure IPAM pools. The structure of `pools` block is documented below.
     */
    pools?: pulumi.Input<pulumi.Input<inputs.sys.IpamPool>[]>;
    /**
     * Configure IPAM allocation rules. The structure of `rules` block is documented below.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.sys.IpamRule>[]>;
    /**
     * Configure the type of IPAM server to use. Valid values: `cloud`, `fabric-root`.
     */
    serverType?: pulumi.Input<string>;
    /**
     * Enable/disable IP address management services. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
