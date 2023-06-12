// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure software switch interfaces by grouping physical and WiFi interfaces.
 *
 * ## Import
 *
 * System SwitchInterface can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:sys/switchinterface:Switchinterface labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:sys/switchinterface:Switchinterface labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Switchinterface extends pulumi.CustomResource {
    /**
     * Get an existing Switchinterface resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SwitchinterfaceState, opts?: pulumi.CustomResourceOptions): Switchinterface {
        return new Switchinterface(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:sys/switchinterface:Switchinterface';

    /**
     * Returns true if the given object is an instance of Switchinterface.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Switchinterface {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Switchinterface.__pulumiType;
    }

    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Allow any traffic between switch interfaces or require firewall policies to allow traffic between switch interfaces. Valid values: `implicit`, `explicit`.
     */
    public readonly intraSwitchPolicy!: pulumi.Output<string>;
    /**
     * Duration for which MAC addresses are held in the ARP table (300 - 8640000 sec, default = 300).
     */
    public readonly macTtl!: pulumi.Output<number>;
    /**
     * Names of the interfaces that belong to the virtual switch. The structure of `member` block is documented below.
     */
    public readonly members!: pulumi.Output<outputs.sys.SwitchinterfaceMember[] | undefined>;
    /**
     * Interface name (name cannot be in use by any other interfaces, VLANs, or inter-VDOM links).
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Enable/disable port spanning. Port spanning echoes traffic received by the software switch to the span destination port. Valid values: `disable`, `enable`.
     */
    public readonly span!: pulumi.Output<string>;
    /**
     * SPAN destination port name. All traffic on the SPAN source ports is echoed to the SPAN destination port.
     */
    public readonly spanDestPort!: pulumi.Output<string>;
    /**
     * The direction in which the SPAN port operates, either: rx, tx, or both. Valid values: `rx`, `tx`, `both`.
     */
    public readonly spanDirection!: pulumi.Output<string>;
    /**
     * Physical interface name. Port spanning echoes all traffic on the SPAN source ports to the SPAN destination port. The structure of `spanSourcePort` block is documented below.
     */
    public readonly spanSourcePorts!: pulumi.Output<outputs.sys.SwitchinterfaceSpanSourcePort[] | undefined>;
    /**
     * Type of switch based on functionality: switch for normal functionality, or hub to duplicate packets to all port members. Valid values: `switch`, `hub`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * VDOM that the software switch belongs to.
     */
    public readonly vdom!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Switchinterface resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SwitchinterfaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SwitchinterfaceArgs | SwitchinterfaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SwitchinterfaceState | undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["intraSwitchPolicy"] = state ? state.intraSwitchPolicy : undefined;
            resourceInputs["macTtl"] = state ? state.macTtl : undefined;
            resourceInputs["members"] = state ? state.members : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["span"] = state ? state.span : undefined;
            resourceInputs["spanDestPort"] = state ? state.spanDestPort : undefined;
            resourceInputs["spanDirection"] = state ? state.spanDirection : undefined;
            resourceInputs["spanSourcePorts"] = state ? state.spanSourcePorts : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["vdom"] = state ? state.vdom : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SwitchinterfaceArgs | undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["intraSwitchPolicy"] = args ? args.intraSwitchPolicy : undefined;
            resourceInputs["macTtl"] = args ? args.macTtl : undefined;
            resourceInputs["members"] = args ? args.members : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["span"] = args ? args.span : undefined;
            resourceInputs["spanDestPort"] = args ? args.spanDestPort : undefined;
            resourceInputs["spanDirection"] = args ? args.spanDirection : undefined;
            resourceInputs["spanSourcePorts"] = args ? args.spanSourcePorts : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["vdom"] = args ? args.vdom : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Switchinterface.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Switchinterface resources.
 */
export interface SwitchinterfaceState {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Allow any traffic between switch interfaces or require firewall policies to allow traffic between switch interfaces. Valid values: `implicit`, `explicit`.
     */
    intraSwitchPolicy?: pulumi.Input<string>;
    /**
     * Duration for which MAC addresses are held in the ARP table (300 - 8640000 sec, default = 300).
     */
    macTtl?: pulumi.Input<number>;
    /**
     * Names of the interfaces that belong to the virtual switch. The structure of `member` block is documented below.
     */
    members?: pulumi.Input<pulumi.Input<inputs.sys.SwitchinterfaceMember>[]>;
    /**
     * Interface name (name cannot be in use by any other interfaces, VLANs, or inter-VDOM links).
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable port spanning. Port spanning echoes traffic received by the software switch to the span destination port. Valid values: `disable`, `enable`.
     */
    span?: pulumi.Input<string>;
    /**
     * SPAN destination port name. All traffic on the SPAN source ports is echoed to the SPAN destination port.
     */
    spanDestPort?: pulumi.Input<string>;
    /**
     * The direction in which the SPAN port operates, either: rx, tx, or both. Valid values: `rx`, `tx`, `both`.
     */
    spanDirection?: pulumi.Input<string>;
    /**
     * Physical interface name. Port spanning echoes all traffic on the SPAN source ports to the SPAN destination port. The structure of `spanSourcePort` block is documented below.
     */
    spanSourcePorts?: pulumi.Input<pulumi.Input<inputs.sys.SwitchinterfaceSpanSourcePort>[]>;
    /**
     * Type of switch based on functionality: switch for normal functionality, or hub to duplicate packets to all port members. Valid values: `switch`, `hub`.
     */
    type?: pulumi.Input<string>;
    /**
     * VDOM that the software switch belongs to.
     */
    vdom?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Switchinterface resource.
 */
export interface SwitchinterfaceArgs {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Allow any traffic between switch interfaces or require firewall policies to allow traffic between switch interfaces. Valid values: `implicit`, `explicit`.
     */
    intraSwitchPolicy?: pulumi.Input<string>;
    /**
     * Duration for which MAC addresses are held in the ARP table (300 - 8640000 sec, default = 300).
     */
    macTtl?: pulumi.Input<number>;
    /**
     * Names of the interfaces that belong to the virtual switch. The structure of `member` block is documented below.
     */
    members?: pulumi.Input<pulumi.Input<inputs.sys.SwitchinterfaceMember>[]>;
    /**
     * Interface name (name cannot be in use by any other interfaces, VLANs, or inter-VDOM links).
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable port spanning. Port spanning echoes traffic received by the software switch to the span destination port. Valid values: `disable`, `enable`.
     */
    span?: pulumi.Input<string>;
    /**
     * SPAN destination port name. All traffic on the SPAN source ports is echoed to the SPAN destination port.
     */
    spanDestPort?: pulumi.Input<string>;
    /**
     * The direction in which the SPAN port operates, either: rx, tx, or both. Valid values: `rx`, `tx`, `both`.
     */
    spanDirection?: pulumi.Input<string>;
    /**
     * Physical interface name. Port spanning echoes all traffic on the SPAN source ports to the SPAN destination port. The structure of `spanSourcePort` block is documented below.
     */
    spanSourcePorts?: pulumi.Input<pulumi.Input<inputs.sys.SwitchinterfaceSpanSourcePort>[]>;
    /**
     * Type of switch based on functionality: switch for normal functionality, or hub to duplicate packets to all port members. Valid values: `switch`, `hub`.
     */
    type?: pulumi.Input<string>;
    /**
     * VDOM that the software switch belongs to.
     */
    vdom?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}