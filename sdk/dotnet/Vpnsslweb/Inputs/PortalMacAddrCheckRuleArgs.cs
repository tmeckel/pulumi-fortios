// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Vpnsslweb.Inputs
{

    public sealed class PortalMacAddrCheckRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("macAddrLists")]
        private InputList<Inputs.PortalMacAddrCheckRuleMacAddrListArgs>? _macAddrLists;

        /// <summary>
        /// Client MAC address list. The structure of `mac_addr_list` block is documented below.
        /// </summary>
        public InputList<Inputs.PortalMacAddrCheckRuleMacAddrListArgs> MacAddrLists
        {
            get => _macAddrLists ?? (_macAddrLists = new InputList<Inputs.PortalMacAddrCheckRuleMacAddrListArgs>());
            set => _macAddrLists = value;
        }

        /// <summary>
        /// Client MAC address mask.
        /// </summary>
        [Input("macAddrMask")]
        public Input<int>? MacAddrMask { get; set; }

        /// <summary>
        /// Client MAC address check rule name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public PortalMacAddrCheckRuleArgs()
        {
        }
        public static new PortalMacAddrCheckRuleArgs Empty => new PortalMacAddrCheckRuleArgs();
    }
}
