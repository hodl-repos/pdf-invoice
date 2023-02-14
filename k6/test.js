import http from "k6/http";

export default function () {
  const url = "http://localhost:12003/v1/generate";
  const payload = JSON.stringify({
    style: {
      localeCode: "de",
      languageCode: "de",
      layout: "DIN_5008B",
      showMarkerPuncher: true,
      showMarkerFolding: true,
      showBankPaymentQrCode: true,
    },
    invoiceAddress: {
      name: "Kunde Mauracher",
      street1: "Straße 1",
      street2: "Top11",
      zip: "1130",
      city: "Wien",
      country: "Austria",
    },
    invoiceInformation: {
      invoiceNumber: "12-30392",
      invoiceDate: "2022-12-23T15:04:05Z",
      dueDate: "2022-12-28T15:04:05Z",
    },
    customerAddress: {
      name: "Auftraggeber Mauracher",
      street1: "Straße 2",
      street2: "Top12",
      zip: "1131",
      city: "Wiefn",
      country: "ÖSTERREICH",
    },
    invoiceData: {
      showNetColumn: true,
      showGrossColumn: true,
      showTaxColumn: true,
      showAmountColumn: true,
      showNetSum: true,
      showTaxSum: true,
      showGrossSum: true,
      sumDiscountPercentage: 10.0,
      sumDiscountFixed: 21.99,
      rows: [
        {
          name: "Anzug",
          description: "3-tlg",
          amount: 1,
          amountUnit: "stk",
          net: 666.66,
          tax: 12.0,
          gross: 888.99,
          discountPercentage: 10.0,
          discountFixed: 15.3,
        },
      ],
    },
    invoiceDataSuffix: "thx for shopping",
    bankPaymentData: {
      accountHolder: "Rotnkopf OG",
      bankName: "VOLKSBANK",
      iban: "AT123456789023",
      bic: "RZTIAT3364",
    },
  });

  const params = {
    headers: {
      "Content-Type": "application/json",
    },
  };

  http.post(url, payload, params);
}
