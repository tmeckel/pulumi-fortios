// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Webfilter.Inputs
{

    public sealed class ProfileWispServerGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Server name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ProfileWispServerGetArgs()
        {
        }
        public static new ProfileWispServerGetArgs Empty => new ProfileWispServerGetArgs();
    }
}
