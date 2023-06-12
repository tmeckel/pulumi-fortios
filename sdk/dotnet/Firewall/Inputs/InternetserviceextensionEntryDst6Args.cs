// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Inputs
{

    public sealed class InternetserviceextensionEntryDst6Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Select the destination address or address group object from available options.
        /// 
        /// The `dst6` block supports:
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public InternetserviceextensionEntryDst6Args()
        {
        }
        public static new InternetserviceextensionEntryDst6Args Empty => new InternetserviceextensionEntryDst6Args();
    }
}
