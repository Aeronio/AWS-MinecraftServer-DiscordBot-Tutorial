//This is used by the Discord Bot to reboot the EC2 instance manually.
//Place this in a Lambda with Javascript runtime.

const AWS = require("aws-sdk");
var ec2 = new AWS.EC2();

exports.handler = async (event) => {
  try {
    var result;
    var params = {
      InstanceIds: [process.env.INSTANCE_ID],
    };
    var data = await ec2.describeInstances(params).promise();
    var instance = data.Reservations[0].Instances[0];

    if (instance.State.Name !== "stopped") {
      result = "instance running";
      console.log("rebooting the instance...");
      var reboot_data = await ec2.rebootInstances(params).promise();
      result = "instance is rebooting";
    } else {
      result = "instance not running";
    }
    const response = {
      statusCode: 200,
      body: result,
    };
    return response;
  } catch (error) {
    console.error(error);
    const response = {
      statusCode: 500,
      body: "error during script",
    };
    return response;
  }
};