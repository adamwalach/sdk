import 'package:test/test.dart';
import 'package:ory_client/ory_client.dart';

// tests for SubmitSelfServiceSettingsFlowWithTotpMethodBody
void main() {
  final instance = SubmitSelfServiceSettingsFlowWithTotpMethodBodyBuilder();
  // TODO add properties to the builder and call build()

  group(SubmitSelfServiceSettingsFlowWithTotpMethodBody, () {
    // CSRFToken is the anti-CSRF token
    // String csrfToken
    test('to test the property `csrfToken`', () async {
      // TODO
    });

    // Method  Should be set to \"totp\" when trying to add, update, or remove a totp pairing.
    // String method
    test('to test the property `method`', () async {
      // TODO
    });

    // ValidationTOTP must contain a valid TOTP based on the
    // String totpCode
    test('to test the property `totpCode`', () async {
      // TODO
    });

    // UnlinkTOTP if true will remove the TOTP pairing, effectively removing the credential. This can be used to set up a new TOTP device.
    // bool totpUnlink
    test('to test the property `totpUnlink`', () async {
      // TODO
    });

  });
}
