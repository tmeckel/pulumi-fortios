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

    public sealed class InterfaceFailAlertInterfaceGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Names of the physical interfaces belonging to the aggregate or redundant interface.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public InterfaceFailAlertInterfaceGetArgs()
        {
        }
        public static new InterfaceFailAlertInterfaceGetArgs Empty => new InterfaceFailAlertInterfaceGetArgs();
    }
}
