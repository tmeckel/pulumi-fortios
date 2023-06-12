// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router.Inputs
{

    public sealed class BgpVrfLeakArgs : global::Pulumi.ResourceArgs
    {
        [Input("targets")]
        private InputList<Inputs.BgpVrfLeakTargetArgs>? _targets;

        /// <summary>
        /// Target VRF table. The structure of `target` block is documented below.
        /// </summary>
        public InputList<Inputs.BgpVrfLeakTargetArgs> Targets
        {
            get => _targets ?? (_targets = new InputList<Inputs.BgpVrfLeakTargetArgs>());
            set => _targets = value;
        }

        /// <summary>
        /// Origin VRF ID &lt;0 - 31&gt;.
        /// </summary>
        [Input("vrf")]
        public Input<string>? Vrf { get; set; }

        public BgpVrfLeakArgs()
        {
        }
        public static new BgpVrfLeakArgs Empty => new BgpVrfLeakArgs();
    }
}