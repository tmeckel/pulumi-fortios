// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Switchcontroller.Outputs
{

    [OutputType]
    public sealed class ManagedswitchStpInstance
    {
        /// <summary>
        /// Instance ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Priority. Valid values: `0`, `4096`, `8192`, `12288`, `16384`, `20480`, `24576`, `28672`, `32768`, `36864`, `40960`, `45056`, `49152`, `53248`, `57344`, `61440`.
        /// </summary>
        public readonly string? Priority;

        [OutputConstructor]
        private ManagedswitchStpInstance(
            string? id,

            string? priority)
        {
            Id = id;
            Priority = priority;
        }
    }
}
