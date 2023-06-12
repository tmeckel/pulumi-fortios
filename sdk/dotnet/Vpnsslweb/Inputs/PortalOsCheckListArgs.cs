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

    public sealed class PortalOsCheckListArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// OS check options. Valid values: `deny`, `allow`, `check-up-to-date`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Latest OS patch level.
        /// </summary>
        [Input("latestPatchLevel")]
        public Input<string>? LatestPatchLevel { get; set; }

        /// <summary>
        /// Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// OS patch level tolerance.
        /// </summary>
        [Input("tolerance")]
        public Input<int>? Tolerance { get; set; }

        public PortalOsCheckListArgs()
        {
        }
        public static new PortalOsCheckListArgs Empty => new PortalOsCheckListArgs();
    }
}
