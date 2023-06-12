// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure Overlay Controller VPN settings. Applies to FortiOS Version `>= 6.2.4`.
 *
 * ## Import
 *
 * Vpn Ocvpn can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:vpn/ocvpn:Ocvpn labelname VpnOcvpn
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:vpn/ocvpn:Ocvpn labelname VpnOcvpn
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Ocvpn extends pulumi.CustomResource {
    /**
     * Get an existing Ocvpn resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OcvpnState, opts?: pulumi.CustomResourceOptions): Ocvpn {
        return new Ocvpn(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:vpn/ocvpn:Ocvpn';

    /**
     * Returns true if the given object is an instance of Ocvpn.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ocvpn {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ocvpn.__pulumiType;
    }

    /**
     * Enable/disable auto-discovery shortcuts. Valid values: `enable`, `disable`.
     */
    public readonly autoDiscovery!: pulumi.Output<string>;
    /**
     * Control deletion of child short-cut tunnels when the parent tunnel goes down. Valid values: `independent`, `dependent`.
     */
    public readonly autoDiscoveryShortcutMode!: pulumi.Output<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable EAP client authentication. Valid values: `enable`, `disable`.
     */
    public readonly eap!: pulumi.Output<string>;
    /**
     * EAP authentication user group.
     */
    public readonly eapUsers!: pulumi.Output<string>;
    /**
     * Configure FortiClient settings. The structure of `forticlientAccess` block is documented below.
     */
    public readonly forticlientAccess!: pulumi.Output<outputs.vpn.OcvpnForticlientAccess>;
    /**
     * Class B subnet reserved for private IP address assignment.
     */
    public readonly ipAllocationBlock!: pulumi.Output<string>;
    /**
     * Enable/disable multipath redundancy. Valid values: `enable`, `disable`.
     */
    public readonly multipath!: pulumi.Output<string>;
    /**
     * Enable/disable inter-overlay source NAT. Valid values: `enable`, `disable`.
     */
    public readonly nat!: pulumi.Output<string>;
    /**
     * Network overlays to register with Overlay Controller VPN service. The structure of `overlays` block is documented below.
     */
    public readonly overlays!: pulumi.Output<outputs.vpn.OcvpnOverlay[] | undefined>;
    /**
     * Overlay Controller VPN polling interval.
     */
    public readonly pollInterval!: pulumi.Output<number>;
    /**
     * Set device role. Valid values: `spoke`, `primary-hub`, `secondary-hub`.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * Enable/disable adding OCVPN tunnels to SDWAN. Valid values: `enable`, `disable`.
     */
    public readonly sdwan!: pulumi.Output<string>;
    /**
     * Set SD-WAN zone.
     */
    public readonly sdwanZone!: pulumi.Output<string>;
    /**
     * Enable/disable Overlay Controller cloud assisted VPN. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * FortiGate WAN interfaces to use with OCVPN. The structure of `wanInterface` block is documented below.
     */
    public readonly wanInterfaces!: pulumi.Output<outputs.vpn.OcvpnWanInterface[] | undefined>;

    /**
     * Create a Ocvpn resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: OcvpnArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OcvpnArgs | OcvpnState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OcvpnState | undefined;
            resourceInputs["autoDiscovery"] = state ? state.autoDiscovery : undefined;
            resourceInputs["autoDiscoveryShortcutMode"] = state ? state.autoDiscoveryShortcutMode : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["eap"] = state ? state.eap : undefined;
            resourceInputs["eapUsers"] = state ? state.eapUsers : undefined;
            resourceInputs["forticlientAccess"] = state ? state.forticlientAccess : undefined;
            resourceInputs["ipAllocationBlock"] = state ? state.ipAllocationBlock : undefined;
            resourceInputs["multipath"] = state ? state.multipath : undefined;
            resourceInputs["nat"] = state ? state.nat : undefined;
            resourceInputs["overlays"] = state ? state.overlays : undefined;
            resourceInputs["pollInterval"] = state ? state.pollInterval : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["sdwan"] = state ? state.sdwan : undefined;
            resourceInputs["sdwanZone"] = state ? state.sdwanZone : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["wanInterfaces"] = state ? state.wanInterfaces : undefined;
        } else {
            const args = argsOrState as OcvpnArgs | undefined;
            resourceInputs["autoDiscovery"] = args ? args.autoDiscovery : undefined;
            resourceInputs["autoDiscoveryShortcutMode"] = args ? args.autoDiscoveryShortcutMode : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["eap"] = args ? args.eap : undefined;
            resourceInputs["eapUsers"] = args ? args.eapUsers : undefined;
            resourceInputs["forticlientAccess"] = args ? args.forticlientAccess : undefined;
            resourceInputs["ipAllocationBlock"] = args ? args.ipAllocationBlock : undefined;
            resourceInputs["multipath"] = args ? args.multipath : undefined;
            resourceInputs["nat"] = args ? args.nat : undefined;
            resourceInputs["overlays"] = args ? args.overlays : undefined;
            resourceInputs["pollInterval"] = args ? args.pollInterval : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["sdwan"] = args ? args.sdwan : undefined;
            resourceInputs["sdwanZone"] = args ? args.sdwanZone : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["wanInterfaces"] = args ? args.wanInterfaces : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Ocvpn.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Ocvpn resources.
 */
export interface OcvpnState {
    /**
     * Enable/disable auto-discovery shortcuts. Valid values: `enable`, `disable`.
     */
    autoDiscovery?: pulumi.Input<string>;
    /**
     * Control deletion of child short-cut tunnels when the parent tunnel goes down. Valid values: `independent`, `dependent`.
     */
    autoDiscoveryShortcutMode?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable EAP client authentication. Valid values: `enable`, `disable`.
     */
    eap?: pulumi.Input<string>;
    /**
     * EAP authentication user group.
     */
    eapUsers?: pulumi.Input<string>;
    /**
     * Configure FortiClient settings. The structure of `forticlientAccess` block is documented below.
     */
    forticlientAccess?: pulumi.Input<inputs.vpn.OcvpnForticlientAccess>;
    /**
     * Class B subnet reserved for private IP address assignment.
     */
    ipAllocationBlock?: pulumi.Input<string>;
    /**
     * Enable/disable multipath redundancy. Valid values: `enable`, `disable`.
     */
    multipath?: pulumi.Input<string>;
    /**
     * Enable/disable inter-overlay source NAT. Valid values: `enable`, `disable`.
     */
    nat?: pulumi.Input<string>;
    /**
     * Network overlays to register with Overlay Controller VPN service. The structure of `overlays` block is documented below.
     */
    overlays?: pulumi.Input<pulumi.Input<inputs.vpn.OcvpnOverlay>[]>;
    /**
     * Overlay Controller VPN polling interval.
     */
    pollInterval?: pulumi.Input<number>;
    /**
     * Set device role. Valid values: `spoke`, `primary-hub`, `secondary-hub`.
     */
    role?: pulumi.Input<string>;
    /**
     * Enable/disable adding OCVPN tunnels to SDWAN. Valid values: `enable`, `disable`.
     */
    sdwan?: pulumi.Input<string>;
    /**
     * Set SD-WAN zone.
     */
    sdwanZone?: pulumi.Input<string>;
    /**
     * Enable/disable Overlay Controller cloud assisted VPN. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * FortiGate WAN interfaces to use with OCVPN. The structure of `wanInterface` block is documented below.
     */
    wanInterfaces?: pulumi.Input<pulumi.Input<inputs.vpn.OcvpnWanInterface>[]>;
}

/**
 * The set of arguments for constructing a Ocvpn resource.
 */
export interface OcvpnArgs {
    /**
     * Enable/disable auto-discovery shortcuts. Valid values: `enable`, `disable`.
     */
    autoDiscovery?: pulumi.Input<string>;
    /**
     * Control deletion of child short-cut tunnels when the parent tunnel goes down. Valid values: `independent`, `dependent`.
     */
    autoDiscoveryShortcutMode?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable EAP client authentication. Valid values: `enable`, `disable`.
     */
    eap?: pulumi.Input<string>;
    /**
     * EAP authentication user group.
     */
    eapUsers?: pulumi.Input<string>;
    /**
     * Configure FortiClient settings. The structure of `forticlientAccess` block is documented below.
     */
    forticlientAccess?: pulumi.Input<inputs.vpn.OcvpnForticlientAccess>;
    /**
     * Class B subnet reserved for private IP address assignment.
     */
    ipAllocationBlock?: pulumi.Input<string>;
    /**
     * Enable/disable multipath redundancy. Valid values: `enable`, `disable`.
     */
    multipath?: pulumi.Input<string>;
    /**
     * Enable/disable inter-overlay source NAT. Valid values: `enable`, `disable`.
     */
    nat?: pulumi.Input<string>;
    /**
     * Network overlays to register with Overlay Controller VPN service. The structure of `overlays` block is documented below.
     */
    overlays?: pulumi.Input<pulumi.Input<inputs.vpn.OcvpnOverlay>[]>;
    /**
     * Overlay Controller VPN polling interval.
     */
    pollInterval?: pulumi.Input<number>;
    /**
     * Set device role. Valid values: `spoke`, `primary-hub`, `secondary-hub`.
     */
    role?: pulumi.Input<string>;
    /**
     * Enable/disable adding OCVPN tunnels to SDWAN. Valid values: `enable`, `disable`.
     */
    sdwan?: pulumi.Input<string>;
    /**
     * Set SD-WAN zone.
     */
    sdwanZone?: pulumi.Input<string>;
    /**
     * Enable/disable Overlay Controller cloud assisted VPN. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * FortiGate WAN interfaces to use with OCVPN. The structure of `wanInterface` block is documented below.
     */
    wanInterfaces?: pulumi.Input<pulumi.Input<inputs.vpn.OcvpnWanInterface>[]>;
}
