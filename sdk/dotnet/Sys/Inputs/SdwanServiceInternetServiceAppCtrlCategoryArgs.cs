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

    public sealed class SdwanServiceInternetServiceAppCtrlCategoryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Application control category ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        public SdwanServiceInternetServiceAppCtrlCategoryArgs()
        {
        }
        public static new SdwanServiceInternetServiceAppCtrlCategoryArgs Empty => new SdwanServiceInternetServiceAppCtrlCategoryArgs();
    }
}
