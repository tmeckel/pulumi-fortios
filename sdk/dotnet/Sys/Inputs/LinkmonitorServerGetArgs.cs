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

    public sealed class LinkmonitorServerGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Server address.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        public LinkmonitorServerGetArgs()
        {
        }
        public static new LinkmonitorServerGetArgs Empty => new LinkmonitorServerGetArgs();
    }
}
