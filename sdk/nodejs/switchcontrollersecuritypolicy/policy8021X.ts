// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure 802.1x MAC Authentication Bypass (MAB) policies.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.switchcontrollersecuritypolicy.Policy8021X("trname", {
 *     authFailVlan: "disable",
 *     authFailVlanid: 0,
 *     eapPassthru: "disable",
 *     framevidApply: "enable",
 *     guestAuthDelay: 30,
 *     guestVlan: "disable",
 *     guestVlanid: 100,
 *     macAuthBypass: "disable",
 *     openAuth: "disable",
 *     policyType: "802.1X",
 *     radiusTimeoutOverwrite: "disable",
 *     securityMode: "802.1X",
 *     userGroups: [{
 *         name: "Guest-group",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * SwitchControllerSecurityPolicy 8021X can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:switchcontrollersecuritypolicy/policy8021X:Policy8021X labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:switchcontrollersecuritypolicy/policy8021X:Policy8021X labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Policy8021X extends pulumi.CustomResource {
    /**
     * Get an existing Policy8021X resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: Policy8021XState, opts?: pulumi.CustomResourceOptions): Policy8021X {
        return new Policy8021X(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:switchcontrollersecuritypolicy/policy8021X:Policy8021X';

    /**
     * Returns true if the given object is an instance of Policy8021X.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Policy8021X {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Policy8021X.__pulumiType;
    }

    /**
     * Enable to allow limited access to clients that cannot authenticate. Valid values: `disable`, `enable`.
     */
    public readonly authFailVlan!: pulumi.Output<string>;
    /**
     * VLAN ID on which authentication failed.
     */
    public readonly authFailVlanId!: pulumi.Output<string>;
    /**
     * VLAN ID on which authentication failed.
     */
    public readonly authFailVlanid!: pulumi.Output<number>;
    /**
     * Authentication server timeout period (3 - 15 sec, default = 3).
     */
    public readonly authserverTimeoutPeriod!: pulumi.Output<number>;
    /**
     * Enable/disable the authentication server timeout VLAN to allow limited access when RADIUS is unavailable.  Valid values: `disable`, `enable`.
     */
    public readonly authserverTimeoutVlan!: pulumi.Output<string>;
    /**
     * Authentication server timeout VLAN name.
     */
    public readonly authserverTimeoutVlanid!: pulumi.Output<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable automatic inclusion of untagged VLANs. Valid values: `disable`, `enable`.
     */
    public readonly eapAutoUntaggedVlans!: pulumi.Output<string>;
    /**
     * Enable/disable EAP pass-through mode, allowing protocols (such as LLDP) to pass through ports for more flexible authentication. Valid values: `disable`, `enable`.
     */
    public readonly eapPassthru!: pulumi.Output<string>;
    /**
     * Enable/disable the capability to apply the EAP/MAB frame VLAN to the port native VLAN. Valid values: `disable`, `enable`.
     */
    public readonly framevidApply!: pulumi.Output<string>;
    /**
     * Guest authentication delay (1 - 900  sec, default = 30).
     */
    public readonly guestAuthDelay!: pulumi.Output<number>;
    /**
     * Enable the guest VLAN feature to allow limited access to non-802.1X-compliant clients. Valid values: `disable`, `enable`.
     */
    public readonly guestVlan!: pulumi.Output<string>;
    /**
     * Guest VLAN name.
     */
    public readonly guestVlanId!: pulumi.Output<string>;
    /**
     * Guest VLAN ID.
     */
    public readonly guestVlanid!: pulumi.Output<number>;
    /**
     * Enable/disable MAB for this policy. Valid values: `disable`, `enable`.
     */
    public readonly macAuthBypass!: pulumi.Output<string>;
    /**
     * Policy name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Enable/disable open authentication for this policy. Valid values: `disable`, `enable`.
     */
    public readonly openAuth!: pulumi.Output<string>;
    /**
     * Policy type. Valid values: `802.1X`.
     */
    public readonly policyType!: pulumi.Output<string>;
    /**
     * Enable to override the global RADIUS session timeout. Valid values: `disable`, `enable`.
     */
    public readonly radiusTimeoutOverwrite!: pulumi.Output<string>;
    /**
     * Port or MAC based 802.1X security mode. Valid values: `802.1X`, `802.1X-mac-based`.
     */
    public readonly securityMode!: pulumi.Output<string>;
    /**
     * Name of user-group to assign to this MAC Authentication Bypass (MAB) policy. The structure of `userGroup` block is documented below.
     */
    public readonly userGroups!: pulumi.Output<outputs.switchcontrollersecuritypolicy.Policy8021XUserGroup[] | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Policy8021X resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: Policy8021XArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: Policy8021XArgs | Policy8021XState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as Policy8021XState | undefined;
            resourceInputs["authFailVlan"] = state ? state.authFailVlan : undefined;
            resourceInputs["authFailVlanId"] = state ? state.authFailVlanId : undefined;
            resourceInputs["authFailVlanid"] = state ? state.authFailVlanid : undefined;
            resourceInputs["authserverTimeoutPeriod"] = state ? state.authserverTimeoutPeriod : undefined;
            resourceInputs["authserverTimeoutVlan"] = state ? state.authserverTimeoutVlan : undefined;
            resourceInputs["authserverTimeoutVlanid"] = state ? state.authserverTimeoutVlanid : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["eapAutoUntaggedVlans"] = state ? state.eapAutoUntaggedVlans : undefined;
            resourceInputs["eapPassthru"] = state ? state.eapPassthru : undefined;
            resourceInputs["framevidApply"] = state ? state.framevidApply : undefined;
            resourceInputs["guestAuthDelay"] = state ? state.guestAuthDelay : undefined;
            resourceInputs["guestVlan"] = state ? state.guestVlan : undefined;
            resourceInputs["guestVlanId"] = state ? state.guestVlanId : undefined;
            resourceInputs["guestVlanid"] = state ? state.guestVlanid : undefined;
            resourceInputs["macAuthBypass"] = state ? state.macAuthBypass : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["openAuth"] = state ? state.openAuth : undefined;
            resourceInputs["policyType"] = state ? state.policyType : undefined;
            resourceInputs["radiusTimeoutOverwrite"] = state ? state.radiusTimeoutOverwrite : undefined;
            resourceInputs["securityMode"] = state ? state.securityMode : undefined;
            resourceInputs["userGroups"] = state ? state.userGroups : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as Policy8021XArgs | undefined;
            resourceInputs["authFailVlan"] = args ? args.authFailVlan : undefined;
            resourceInputs["authFailVlanId"] = args ? args.authFailVlanId : undefined;
            resourceInputs["authFailVlanid"] = args ? args.authFailVlanid : undefined;
            resourceInputs["authserverTimeoutPeriod"] = args ? args.authserverTimeoutPeriod : undefined;
            resourceInputs["authserverTimeoutVlan"] = args ? args.authserverTimeoutVlan : undefined;
            resourceInputs["authserverTimeoutVlanid"] = args ? args.authserverTimeoutVlanid : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["eapAutoUntaggedVlans"] = args ? args.eapAutoUntaggedVlans : undefined;
            resourceInputs["eapPassthru"] = args ? args.eapPassthru : undefined;
            resourceInputs["framevidApply"] = args ? args.framevidApply : undefined;
            resourceInputs["guestAuthDelay"] = args ? args.guestAuthDelay : undefined;
            resourceInputs["guestVlan"] = args ? args.guestVlan : undefined;
            resourceInputs["guestVlanId"] = args ? args.guestVlanId : undefined;
            resourceInputs["guestVlanid"] = args ? args.guestVlanid : undefined;
            resourceInputs["macAuthBypass"] = args ? args.macAuthBypass : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["openAuth"] = args ? args.openAuth : undefined;
            resourceInputs["policyType"] = args ? args.policyType : undefined;
            resourceInputs["radiusTimeoutOverwrite"] = args ? args.radiusTimeoutOverwrite : undefined;
            resourceInputs["securityMode"] = args ? args.securityMode : undefined;
            resourceInputs["userGroups"] = args ? args.userGroups : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Policy8021X.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Policy8021X resources.
 */
export interface Policy8021XState {
    /**
     * Enable to allow limited access to clients that cannot authenticate. Valid values: `disable`, `enable`.
     */
    authFailVlan?: pulumi.Input<string>;
    /**
     * VLAN ID on which authentication failed.
     */
    authFailVlanId?: pulumi.Input<string>;
    /**
     * VLAN ID on which authentication failed.
     */
    authFailVlanid?: pulumi.Input<number>;
    /**
     * Authentication server timeout period (3 - 15 sec, default = 3).
     */
    authserverTimeoutPeriod?: pulumi.Input<number>;
    /**
     * Enable/disable the authentication server timeout VLAN to allow limited access when RADIUS is unavailable.  Valid values: `disable`, `enable`.
     */
    authserverTimeoutVlan?: pulumi.Input<string>;
    /**
     * Authentication server timeout VLAN name.
     */
    authserverTimeoutVlanid?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable automatic inclusion of untagged VLANs. Valid values: `disable`, `enable`.
     */
    eapAutoUntaggedVlans?: pulumi.Input<string>;
    /**
     * Enable/disable EAP pass-through mode, allowing protocols (such as LLDP) to pass through ports for more flexible authentication. Valid values: `disable`, `enable`.
     */
    eapPassthru?: pulumi.Input<string>;
    /**
     * Enable/disable the capability to apply the EAP/MAB frame VLAN to the port native VLAN. Valid values: `disable`, `enable`.
     */
    framevidApply?: pulumi.Input<string>;
    /**
     * Guest authentication delay (1 - 900  sec, default = 30).
     */
    guestAuthDelay?: pulumi.Input<number>;
    /**
     * Enable the guest VLAN feature to allow limited access to non-802.1X-compliant clients. Valid values: `disable`, `enable`.
     */
    guestVlan?: pulumi.Input<string>;
    /**
     * Guest VLAN name.
     */
    guestVlanId?: pulumi.Input<string>;
    /**
     * Guest VLAN ID.
     */
    guestVlanid?: pulumi.Input<number>;
    /**
     * Enable/disable MAB for this policy. Valid values: `disable`, `enable`.
     */
    macAuthBypass?: pulumi.Input<string>;
    /**
     * Policy name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable open authentication for this policy. Valid values: `disable`, `enable`.
     */
    openAuth?: pulumi.Input<string>;
    /**
     * Policy type. Valid values: `802.1X`.
     */
    policyType?: pulumi.Input<string>;
    /**
     * Enable to override the global RADIUS session timeout. Valid values: `disable`, `enable`.
     */
    radiusTimeoutOverwrite?: pulumi.Input<string>;
    /**
     * Port or MAC based 802.1X security mode. Valid values: `802.1X`, `802.1X-mac-based`.
     */
    securityMode?: pulumi.Input<string>;
    /**
     * Name of user-group to assign to this MAC Authentication Bypass (MAB) policy. The structure of `userGroup` block is documented below.
     */
    userGroups?: pulumi.Input<pulumi.Input<inputs.switchcontrollersecuritypolicy.Policy8021XUserGroup>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Policy8021X resource.
 */
export interface Policy8021XArgs {
    /**
     * Enable to allow limited access to clients that cannot authenticate. Valid values: `disable`, `enable`.
     */
    authFailVlan?: pulumi.Input<string>;
    /**
     * VLAN ID on which authentication failed.
     */
    authFailVlanId?: pulumi.Input<string>;
    /**
     * VLAN ID on which authentication failed.
     */
    authFailVlanid?: pulumi.Input<number>;
    /**
     * Authentication server timeout period (3 - 15 sec, default = 3).
     */
    authserverTimeoutPeriod?: pulumi.Input<number>;
    /**
     * Enable/disable the authentication server timeout VLAN to allow limited access when RADIUS is unavailable.  Valid values: `disable`, `enable`.
     */
    authserverTimeoutVlan?: pulumi.Input<string>;
    /**
     * Authentication server timeout VLAN name.
     */
    authserverTimeoutVlanid?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable automatic inclusion of untagged VLANs. Valid values: `disable`, `enable`.
     */
    eapAutoUntaggedVlans?: pulumi.Input<string>;
    /**
     * Enable/disable EAP pass-through mode, allowing protocols (such as LLDP) to pass through ports for more flexible authentication. Valid values: `disable`, `enable`.
     */
    eapPassthru?: pulumi.Input<string>;
    /**
     * Enable/disable the capability to apply the EAP/MAB frame VLAN to the port native VLAN. Valid values: `disable`, `enable`.
     */
    framevidApply?: pulumi.Input<string>;
    /**
     * Guest authentication delay (1 - 900  sec, default = 30).
     */
    guestAuthDelay?: pulumi.Input<number>;
    /**
     * Enable the guest VLAN feature to allow limited access to non-802.1X-compliant clients. Valid values: `disable`, `enable`.
     */
    guestVlan?: pulumi.Input<string>;
    /**
     * Guest VLAN name.
     */
    guestVlanId?: pulumi.Input<string>;
    /**
     * Guest VLAN ID.
     */
    guestVlanid?: pulumi.Input<number>;
    /**
     * Enable/disable MAB for this policy. Valid values: `disable`, `enable`.
     */
    macAuthBypass?: pulumi.Input<string>;
    /**
     * Policy name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable open authentication for this policy. Valid values: `disable`, `enable`.
     */
    openAuth?: pulumi.Input<string>;
    /**
     * Policy type. Valid values: `802.1X`.
     */
    policyType?: pulumi.Input<string>;
    /**
     * Enable to override the global RADIUS session timeout. Valid values: `disable`, `enable`.
     */
    radiusTimeoutOverwrite?: pulumi.Input<string>;
    /**
     * Port or MAC based 802.1X security mode. Valid values: `802.1X`, `802.1X-mac-based`.
     */
    securityMode?: pulumi.Input<string>;
    /**
     * Name of user-group to assign to this MAC Authentication Bypass (MAB) policy. The structure of `userGroup` block is documented below.
     */
    userGroups?: pulumi.Input<pulumi.Input<inputs.switchcontrollersecuritypolicy.Policy8021XUserGroup>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
