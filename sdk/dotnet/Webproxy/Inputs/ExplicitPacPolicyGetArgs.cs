// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Webproxy.Inputs
{

    public sealed class ExplicitPacPolicyGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional comments.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        [Input("dstaddrs")]
        private InputList<Inputs.ExplicitPacPolicyDstaddrGetArgs>? _dstaddrs;

        /// <summary>
        /// Destination address objects. The structure of `dstaddr` block is documented below.
        /// </summary>
        public InputList<Inputs.ExplicitPacPolicyDstaddrGetArgs> Dstaddrs
        {
            get => _dstaddrs ?? (_dstaddrs = new InputList<Inputs.ExplicitPacPolicyDstaddrGetArgs>());
            set => _dstaddrs = value;
        }

        /// <summary>
        /// PAC file contents enclosed in quotes (maximum of 256K bytes).
        /// </summary>
        [Input("pacFileData")]
        public Input<string>? PacFileData { get; set; }

        /// <summary>
        /// Pac file name.
        /// </summary>
        [Input("pacFileName")]
        public Input<string>? PacFileName { get; set; }

        /// <summary>
        /// Policy ID.
        /// </summary>
        [Input("policyid")]
        public Input<int>? Policyid { get; set; }

        [Input("srcaddr6s")]
        private InputList<Inputs.ExplicitPacPolicySrcaddr6GetArgs>? _srcaddr6s;

        /// <summary>
        /// Source address6 objects. The structure of `srcaddr6` block is documented below.
        /// </summary>
        public InputList<Inputs.ExplicitPacPolicySrcaddr6GetArgs> Srcaddr6s
        {
            get => _srcaddr6s ?? (_srcaddr6s = new InputList<Inputs.ExplicitPacPolicySrcaddr6GetArgs>());
            set => _srcaddr6s = value;
        }

        [Input("srcaddrs")]
        private InputList<Inputs.ExplicitPacPolicySrcaddrGetArgs>? _srcaddrs;

        /// <summary>
        /// Source address objects. The structure of `srcaddr` block is documented below.
        /// </summary>
        public InputList<Inputs.ExplicitPacPolicySrcaddrGetArgs> Srcaddrs
        {
            get => _srcaddrs ?? (_srcaddrs = new InputList<Inputs.ExplicitPacPolicySrcaddrGetArgs>());
            set => _srcaddrs = value;
        }

        /// <summary>
        /// Enable/disable policy. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public ExplicitPacPolicyGetArgs()
        {
        }
        public static new ExplicitPacPolicyGetArgs Empty => new ExplicitPacPolicyGetArgs();
    }
}