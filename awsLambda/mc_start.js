//This is used by the Discord Bot to start the EC2 instance manually.
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
      var launch_time = new Date(instance.LaunchTime);
      var today = new Date();
      result = "instance running, started at " + launch_time;
    } else {
      var start_data = await ec2.startInstances(params).promise();
      result = "instance started";
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