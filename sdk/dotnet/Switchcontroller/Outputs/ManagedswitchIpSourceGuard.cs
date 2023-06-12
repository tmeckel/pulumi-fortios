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
    public sealed class ManagedswitchIpSourceGuard
    {
        /// <summary>
        /// IP and MAC address configuration. The structure of `binding_entry` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ManagedswitchIpSourceGuardBindingEntry> BindingEntries;
        /// <summary>
        /// Description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Ingress interface to which source guard is bound.
        /// </summary>
        public readonly string? Port;

        [OutputConstructor]
        private ManagedswitchIpSourceGuard(
            ImmutableArray<Outputs.ManagedswitchIpSourceGuardBindingEntry> bindingEntries,

            string? description,

            string? port)
        {
            BindingEntries = bindingEntries;
            Description = description;
            Port = port;
        }
    }
}
