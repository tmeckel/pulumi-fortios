// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Extensioncontroller.Inputs
{

    public sealed class ExtenderprofileCellularDataplanGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Dataplan name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ExtenderprofileCellularDataplanGetArgs()
        {
        }
        public static new ExtenderprofileCellularDataplanGetArgs Empty => new ExtenderprofileCellularDataplanGetArgs();
    }
}
