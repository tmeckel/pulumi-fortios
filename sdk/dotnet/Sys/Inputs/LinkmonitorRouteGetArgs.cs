// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Sys.Inputs
{

    public sealed class LinkmonitorRouteGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP and netmask (x.x.x.x/y).
        /// </summary>
        [Input("subnet")]
        public Input<string>? Subnet { get; set; }

        public LinkmonitorRouteGetArgs()
        {
        }
        public static new LinkmonitorRouteGetArgs Empty => new LinkmonitorRouteGetArgs();
    }
}
