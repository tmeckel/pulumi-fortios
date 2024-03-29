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

    public sealed class BgpVrf6Args : global::Pulumi.ResourceArgs
    {
        [Input("leakTargets")]
        private InputList<Inputs.BgpVrf6LeakTargetArgs>? _leakTargets;

        /// <summary>
        /// Target VRF table. The structure of `leak_target` block is documented below.
        /// </summary>
        public InputList<Inputs.BgpVrf6LeakTargetArgs> LeakTargets
        {
            get => _leakTargets ?? (_leakTargets = new InputList<Inputs.BgpVrf6LeakTargetArgs>());
            set => _leakTargets = value;
        }

        /// <summary>
        /// BGP VRF leaking table. The structure of `vrf` block is documented below.
        /// </summary>
        [Input("vrf")]
        public Input<string>? Vrf { get; set; }

        public BgpVrf6Args()
        {
        }
        public static new BgpVrf6Args Empty => new BgpVrf6Args();
    }
}
