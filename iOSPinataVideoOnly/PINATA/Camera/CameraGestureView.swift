////
////  CameraGestures.swift
////  PINATA
////
////  Created by Karim Abedrabbo on 11/19/19.
////  Copyright © 2019 Karim Abedrabbo. All rights reserved.
////
//
//import Foundation
//import UIKit
//
//public class CameraGestureView: UIView, UIGestureRecognizerDelegate {
// 
//  internal var longPressStartPoint: CGPoint = .zero
//
//  public weak var delegate: CameraLogicManagerRepresentable
//  
//  init(delegate: CameraLogicManagerRepresentable) {
//    super.init(frame: .zero)
//    
//    self.delegate = delegate
//    
//    let focusTapGestureRecognizer = UITapGestureRecognizer(target: self, action: #selector(handleFocusTapGestureRecognizer(_:)))
//       focusTapGestureRecognizer.delegate = self
//       focusTapGestureRecognizer.numberOfTapsRequired = 1
//       self.addGestureRecognizer(focusTapGestureRecognizer)
//       
//       
//       let panGestureRecognizer = UIPanGestureRecognizer(target: self, action: #selector(handlePanGestureRecognizer(_:)))
//       panGestureRecognizer.delegate = self
//       self.addGestureRecognizer(panGestureRecognizer)
//  
//  }
//  
//  required init?(coder: NSCoder) {
//    fatalError("init(coder:) has not been implemented")
//  }
//  
//     
//   //tap
//   @objc internal func handleRecordTapRecognizer(_ gestureRecognizer: UIGestureRecognizer) {
//       gestureRecognizer.isEnabled = false
//       self.recorder.capturePhoto { [weak self] (error, image) in
//           if error != nil {
//               print("Could not capture photo with error: \(error)")
//           } else if image == nil {
//               print("Image captured is nil")
//           } else {
//               self?.present(UINavigationController(rootViewController: VideoPlayerViewController(image: image!, style: "cameraPhoto")), animated: true, completion: nil)
//
//           }
//       }
//     }
//     
//     //long press
//     @objc internal func handleRecordLongPressRecognizer(_ gestureRecognizer: UIGestureRecognizer) {
//             switch gestureRecognizer.state {
//             case .began:
//                 self.startCapture()
//                 self.longPressStartPoint = gestureRecognizer.location(in: self.previewView)
//                 self.startZoom = CGFloat(self.recorder.videoZoomFactor)
//                 break
//             case .changed:
//                 let newGesturePoint = gestureRecognizer.location(in: self.previewView)
//                 
//                 recordView.center = newGesturePoint
//                 
//                 
//                 let deltaY = self.longPressStartPoint.y / newGesturePoint.y
//                 var zoomScale: CGFloat = 1.0
//                 if deltaY > 1.0 {
//                     zoomScale = pow(deltaY, 1.5)
//                 } else {
//                     zoomScale = pow(deltaY, 10.0)
//                 }
//                 let newZoom = (zoomScale * self.startZoom)
//                 self.recorder.videoZoomFactor = min(max(1.0, CGFloat(newZoom)), CGFloat(self.recorder.maxVideoZoomFactor / 4))
//                 break
//             case .ended:
//                 fallthrough
//             case .cancelled:
//                 fallthrough
//             case .failed:
//                 self.pauseCapture(completion: nil)
//                 
//                 fallthrough
//             default:
//                 break
//             }
//     }
//     
//     
//     @objc internal func handleDeleteLongPressGestureRecognizer(_ gestureRecognizer: UIGestureRecognizer) {
//      
//       
//       let alert = UIAlertController(title: "Delete All?", message: "", preferredStyle: .actionSheet)
//  
//       alert.addAction(UIAlertAction(title: "Keep", style: .default, handler: nil))
//       alert.addAction(UIAlertAction(title: "Delete All", style: .destructive) { [weak self] (alertAction) in
//             self?.handleDeleteAllAction()
//         })
//
//
//       self.present(alert, animated: true)
//       
//     }
//     
//     //tap
//     @objc internal func handleFocusTapGestureRecognizer(_ gestureRecognizer: UIGestureRecognizer) {
//         if !recorder.isAdjustingFocus{
//         let tapPoint = gestureRecognizer.location(in: self.previewView)
//         let focusSize = CGSize(width: 30, height: 30)
//         let focusRect = CGRect(x: tapPoint.x - focusSize.width/2, y: tapPoint.y - focusSize.height/2, width: focusSize.width, height: focusSize.height)
//         self.focusView = FocusIndicatorView(frame: focusRect)
//         self.view.addSubview(focusView!)
//         
//         let recorderPoint = recorder.convertToPointOfInterest(fromViewCoordinates: tapPoint)
//         recorder.autoFocus(at: recorderPoint)
//         }
//     }
//     
//     @objc internal func handlePanGestureRecognizer(_ gestureRecognizer: UIGestureRecognizer) {
//         switch gestureRecognizer.state {
//         case .began:
//             self.longPressStartPoint = gestureRecognizer.location(in: self.previewView)
//             self.startZoom = CGFloat(self.recorder.videoZoomFactor)
//             break
//         case .changed:
//             let newGesturePoint = gestureRecognizer.location(in: self.previewView)
//             
//             
//             
//             let deltaY = self.longPressStartPoint.y / newGesturePoint.y
//             var zoomScale: CGFloat = 1.0
//             if deltaY > 1.0 {
//                 zoomScale = pow(deltaY, 1.5)
//             } else {
//                 zoomScale = pow(deltaY, 3.5)
//             }
//             let newZoom = (zoomScale * self.startZoom)
//             self.recorder.videoZoomFactor = min(max(1.0, CGFloat(newZoom)), CGFloat(self.recorder.maxVideoZoomFactor / 4))
//             break
//         case .ended:
//             fallthrough
//         case .cancelled:
//             fallthrough
//         case .failed:
//             fallthrough
//         default:
//             break
//         }
//     }
//  
//}
//extension CameraViewController: UIGestureRecognizerDelegate {
//
//    
//    
//}
