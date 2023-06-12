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

    public sealed class ExtenderprofileCellularSmsNotificationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// SMS alert list. The structure of `alert` block is documented below.
        /// </summary>
        [Input("alert")]
        public Input<Inputs.ExtenderprofileCellularSmsNotificationAlertArgs>? Alert { get; set; }

        [Input("receivers")]
        private InputList<Inputs.ExtenderprofileCellularSmsNotificationReceiverArgs>? _receivers;

        /// <summary>
        /// SMS notification receiver list. The structure of `receiver` block is documented below.
        /// </summary>
        public InputList<Inputs.ExtenderprofileCellularSmsNotificationReceiverArgs> Receivers
        {
            get => _receivers ?? (_receivers = new InputList<Inputs.ExtenderprofileCellularSmsNotificationReceiverArgs>());
            set => _receivers = value;
        }

        /// <summary>
        /// FortiExtender SMS notification status. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public ExtenderprofileCellularSmsNotificationArgs()
        {
        }
        public static new ExtenderprofileCellularSmsNotificationArgs Empty => new ExtenderprofileCellularSmsNotificationArgs();
    }
}
