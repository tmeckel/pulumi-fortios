// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.User.Inputs
{

    public sealed class SecurityexemptlistRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("devices")]
        private InputList<Inputs.SecurityexemptlistRuleDeviceArgs>? _devices;

        /// <summary>
        /// Devices or device groups. The structure of `devices` block is documented below.
        /// </summary>
        public InputList<Inputs.SecurityexemptlistRuleDeviceArgs> Devices
        {
            get => _devices ?? (_devices = new InputList<Inputs.SecurityexemptlistRuleDeviceArgs>());
            set => _devices = value;
        }

        [Input("dstaddrs")]
        private InputList<Inputs.SecurityexemptlistRuleDstaddrArgs>? _dstaddrs;

        /// <summary>
        /// Destination addresses or address groups. The structure of `dstaddr` block is documented below.
        /// </summary>
        public InputList<Inputs.SecurityexemptlistRuleDstaddrArgs> Dstaddrs
        {
            get => _dstaddrs ?? (_dstaddrs = new InputList<Inputs.SecurityexemptlistRuleDstaddrArgs>());
            set => _dstaddrs = value;
        }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        [Input("services")]
        private InputList<Inputs.SecurityexemptlistRuleServiceArgs>? _services;

        /// <summary>
        /// Destination services. The structure of `service` block is documented below.
        /// </summary>
        public InputList<Inputs.SecurityexemptlistRuleServiceArgs> Services
        {
            get => _services ?? (_services = new InputList<Inputs.SecurityexemptlistRuleServiceArgs>());
            set => _services = value;
        }

        [Input("srcaddrs")]
        private InputList<Inputs.SecurityexemptlistRuleSrcaddrArgs>? _srcaddrs;

        /// <summary>
        /// Source addresses or address groups. The structure of `srcaddr` block is documented below.
        /// </summary>
        public InputList<Inputs.SecurityexemptlistRuleSrcaddrArgs> Srcaddrs
        {
            get => _srcaddrs ?? (_srcaddrs = new InputList<Inputs.SecurityexemptlistRuleSrcaddrArgs>());
            set => _srcaddrs = value;
        }

        public SecurityexemptlistRuleArgs()
        {
        }
        public static new SecurityexemptlistRuleArgs Empty => new SecurityexemptlistRuleArgs();
    }
}