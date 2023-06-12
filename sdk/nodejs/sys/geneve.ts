// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure GENEVE devices. Applies to FortiOS Version `>= 6.2.4`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.sys.Geneve("trname", {
 *     dstport: 22,
 *     "interface": "port2",
 *     ipVersion: "ipv4-unicast",
 *     remoteIp: "1.1.1.1",
 *     remoteIp6: "::",
 *     vni: 0,
 * });
 * ```
 *
 * ## Import
 *
 * System Geneve can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:sys/geneve:Geneve labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:sys/geneve:Geneve labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Geneve extends pulumi.CustomResource {
    /**
     * Get an existing Geneve resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GeneveState, opts?: pulumi.CustomResourceOptions): Geneve {
        return new Geneve(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:sys/geneve:Geneve';

    /**
     * Returns true if the given object is an instance of Geneve.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Geneve {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Geneve.__pulumiType;
    }

    /**
     * GENEVE destination port (1 - 65535, default = 6081).
     */
    public readonly dstport!: pulumi.Output<number>;
    /**
     * Outgoing interface for GENEVE encapsulated traffic.
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * IP version to use for the GENEVE interface and so for communication over the GENEVE. IPv4 or IPv6 unicast. Valid values: `ipv4-unicast`, `ipv6-unicast`.
     */
    public readonly ipVersion!: pulumi.Output<string>;
    /**
     * GENEVE device or interface name. Must be an unique interface name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * IPv4 address of the GENEVE interface on the device at the remote end of the GENEVE.
     */
    public readonly remoteIp!: pulumi.Output<string>;
    /**
     * IPv6 IP address of the GENEVE interface on the device at the remote end of the GENEVE.
     */
    public readonly remoteIp6!: pulumi.Output<string>;
    /**
     * GENEVE type. Valid values: `ethernet`, `ppp`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * GENEVE network ID.
     */
    public readonly vni!: pulumi.Output<number>;

    /**
     * Create a Geneve resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GeneveArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GeneveArgs | GeneveState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GeneveState | undefined;
            resourceInputs["dstport"] = state ? state.dstport : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["ipVersion"] = state ? state.ipVersion : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["remoteIp"] = state ? state.remoteIp : undefined;
            resourceInputs["remoteIp6"] = state ? state.remoteIp6 : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["vni"] = state ? state.vni : undefined;
        } else {
            const args = argsOrState as GeneveArgs | undefined;
            if ((!args || args.interface === undefined) && !opts.urn) {
                throw new Error("Missing required property 'interface'");
            }
            if ((!args || args.ipVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipVersion'");
            }
            if ((!args || args.remoteIp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'remoteIp'");
            }
            if ((!args || args.vni === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vni'");
            }
            resourceInputs["dstport"] = args ? args.dstport : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["ipVersion"] = args ? args.ipVersion : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["remoteIp"] = args ? args.remoteIp : undefined;
            resourceInputs["remoteIp6"] = args ? args.remoteIp6 : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["vni"] = args ? args.vni : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Geneve.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Geneve resources.
 */
export interface GeneveState {
    /**
     * GENEVE destination port (1 - 65535, default = 6081).
     */
    dstport?: pulumi.Input<number>;
    /**
     * Outgoing interface for GENEVE encapsulated traffic.
     */
    interface?: pulumi.Input<string>;
    /**
     * IP version to use for the GENEVE interface and so for communication over the GENEVE. IPv4 or IPv6 unicast. Valid values: `ipv4-unicast`, `ipv6-unicast`.
     */
    ipVersion?: pulumi.Input<string>;
    /**
     * GENEVE device or interface name. Must be an unique interface name.
     */
    name?: pulumi.Input<string>;
    /**
     * IPv4 address of the GENEVE interface on the device at the remote end of the GENEVE.
     */
    remoteIp?: pulumi.Input<string>;
    /**
     * IPv6 IP address of the GENEVE interface on the device at the remote end of the GENEVE.
     */
    remoteIp6?: pulumi.Input<string>;
    /**
     * GENEVE type. Valid values: `ethernet`, `ppp`.
     */
    type?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * GENEVE network ID.
     */
    vni?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Geneve resource.
 */
export interface GeneveArgs {
    /**
     * GENEVE destination port (1 - 65535, default = 6081).
     */
    dstport?: pulumi.Input<number>;
    /**
     * Outgoing interface for GENEVE encapsulated traffic.
     */
    interface: pulumi.Input<string>;
    /**
     * IP version to use for the GENEVE interface and so for communication over the GENEVE. IPv4 or IPv6 unicast. Valid values: `ipv4-unicast`, `ipv6-unicast`.
     */
    ipVersion: pulumi.Input<string>;
    /**
     * GENEVE device or interface name. Must be an unique interface name.
     */
    name?: pulumi.Input<string>;
    /**
     * IPv4 address of the GENEVE interface on the device at the remote end of the GENEVE.
     */
    remoteIp: pulumi.Input<string>;
    /**
     * IPv6 IP address of the GENEVE interface on the device at the remote end of the GENEVE.
     */
    remoteIp6?: pulumi.Input<string>;
    /**
     * GENEVE type. Valid values: `ethernet`, `ppp`.
     */
    type?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * GENEVE network ID.
     */
    vni: pulumi.Input<number>;
}
